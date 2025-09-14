package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sumanchapai/gw/ctemplates"
	"github.com/sumanchapai/gw/env"
	"github.com/sumanchapai/gw/git"
	"github.com/sumanchapai/gw/utils"

	"html/template"
)

func Root(w http.ResponseWriter, r *http.Request) {
	repo := env.GW_REPO()
	if repo == "" {
		fmt.Fprintf(w, "GW_REPO environment variable is empty.")
		return
	}

	// Check if the GW_REPO exists
	if err := utils.FolderExists(repo); err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%s", fmt.Sprintf("GW_REPO: %v. error: %v", repo, err))
		return
	}

	// Check if the GW_REPO is a valid Git repo
	if err := git.IsGitRepo(repo); err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%s", fmt.Sprintf("GW_REPO: %v. error: %v", repo, err))
		return
	}

	// Parse template
	path, err := ctemplates.Path("root.html")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "error getting template path %s", err)
	}
	tmpl := template.Must(template.ParseFiles(path))

	tmplData, err := ctemplates.GetRootData()
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "%s", fmt.Sprintf("Error getting template data. error: %v", err))
		return
	}
	tmpl.Execute(w, tmplData)
}
