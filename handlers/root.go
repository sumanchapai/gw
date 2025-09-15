package handlers

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/sumanchapai/gw/ctemplates"
	"github.com/sumanchapai/gw/env"
	"github.com/sumanchapai/gw/git"
	"github.com/sumanchapai/gw/utils"

	"html/template"
)

func Root(templateFS embed.FS) http.HandlerFunc {

	tmpl := template.Must(template.ParseFS(templateFS, "templates/*.html"))

	return func(w http.ResponseWriter, r *http.Request) {
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

		tmplData, err := ctemplates.GetRootData()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "%s", fmt.Sprintf("Error getting template data. error: %v", err))
			return
		}
		tmpl.Execute(w, tmplData)
	}
}
