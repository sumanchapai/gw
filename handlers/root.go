package handlers

import (
	"fmt"
	"net/http"

	"github.com/sumanchapai/gw/env"
)

func Root(w http.ResponseWriter, r *http.Request) {
	if v := env.GW_REPO(); v == "" {
		fmt.Fprintf(w, "GW_REPO environment variable is empty.")
		return
	}
	fmt.Fprintf(w, "Hello, world!")
}
