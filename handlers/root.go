package handlers

import (
	"fmt"
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
		fmt.Fprintf(w, "%s", fmt.Sprintf("GW_REPO: %v. error: %v", repo, err))
		return
	}

	// Check if the GW_REPO is a valid Git repo
	if err := git.IsGitRepo(repo); err != nil {
		fmt.Fprintf(w, "%s", fmt.Sprintf("GW_REPO: %v. error: %v", repo, err))
		return
	}

	// Parse template
	path, err := ctemplates.Path("root.html")
	if err != nil {
		fmt.Fprintf(w, "error getting template path %s", err)
	}
	tmpl := template.Must(template.ParseFiles(path))
	tmpl.Execute(w, nil)
}
