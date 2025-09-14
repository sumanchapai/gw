package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/sumanchapai/gw/env"
	"github.com/sumanchapai/gw/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	http.HandleFunc("/", handlers.Root)

	// Start server
	hostPort := fmt.Sprintf("%v:%v", env.Host(), env.Port())
	log.Printf("starting server at %v", hostPort)
	log.Fatal(http.ListenAndServe(hostPort, nil))
}
