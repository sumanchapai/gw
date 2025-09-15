package handlers

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/websocket"
	"github.com/sumanchapai/gw/env"
	"github.com/sumanchapai/gw/scripts"
)

type ActionRequest struct {
	Action string   `json:"action"`
	Args   []string `json:"commitmsg"`
}

var upgrader = websocket.Upgrader{}

// GitAction handler streams the result of running the provided git-action
func GitAction(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	defer conn.Close()
	var req ActionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("Invalid JSON"))
		return
	}

	var script string

	// Get the appropriate script name
	switch req.Action {
	case "commit":
		script, err = scripts.Path("commit.sh")
		if err != nil {
			log.Println("Error loading script", script, "err:", err)
			conn.WriteMessage(websocket.TextMessage, []byte("Error loading script"))
			return
		}
	default:
		log.Println("Invalid action received", req.Action)
		conn.WriteMessage(websocket.TextMessage, []byte("Invalid action received"))
		return
	}

	// Run the script appropriately
	cmd := exec.Command(script, req.Args...)
	cmd.Dir = env.GW_REPO()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("Error getting stdout", err)
		conn.WriteMessage(websocket.TextMessage, []byte("Error getting stdout"))
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("Error getting stderr", err)
		conn.WriteMessage(websocket.TextMessage, []byte("Error getting stderr"))
		return
	}

	if err := cmd.Start(); err != nil {
		log.Println("Error starting command", err)
		conn.WriteMessage(websocket.TextMessage, []byte("Error starting command"))
	}

	scanner := bufio.NewScanner(io.MultiReader(stdout, stderr))
	for scanner.Scan() {
		line := scanner.Text()
		conn.WriteMessage(websocket.TextMessage, []byte(line))
	}
	cmd.Wait()
	conn.WriteMessage(websocket.TextMessage, []byte("DONE"))
}
