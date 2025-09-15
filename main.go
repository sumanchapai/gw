package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/sumanchapai/gw/env"
	"github.com/sumanchapai/gw/handlers"
)

// Embed the entire static folder
//
//go:embed static/*
var staticFS embed.FS

// Embed templates
//
//go:embed templates/*
var templatesFS embed.FS

// Embed scripts
//
//go:embed scripts/*
var scriptsFS embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	http.HandleFunc("/", handlers.Root(templatesFS))
	http.HandleFunc("/git-command", handlers.GitCommand)
	http.HandleFunc("/git-action", handlers.GitAction)

	// Create a sub-FS rooted at "static"
	subFS, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(subFS))))

	// Start server
	hostPort := fmt.Sprintf("%v:%v", env.Host(), env.Port())
	log.Printf("starting server at %v", hostPort)
	log.Fatal(http.ListenAndServe(hostPort, nil))
}
