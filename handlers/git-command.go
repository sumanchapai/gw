package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sumanchapai/gw/env"
	"github.com/sumanchapai/gw/git"
)

type GitRequest struct {
	Args []string `json:"args"`
}

type GitResponse struct {
	Success bool   `json:"success"`
	Output  string `json:"output"`
	Error   string `json:"error,omitempty"`
}

// GitCommand is a REST API endpoint that receives git commands from the user
// and returns the result.
func GitCommand(w http.ResponseWriter, r *http.Request) {
	// Parse JSON body
	var req GitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if len(req.Args) == 0 {
		http.Error(w, "No git command provided", http.StatusBadRequest)
		return
	}

	// Run git command
	res := git.SafeRun(env.GW_REPO(), req.Args...)

	// Build response
	resp := GitResponse{
		Success: res.Err == nil,
		Output:  res.Result}

	// Return JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
