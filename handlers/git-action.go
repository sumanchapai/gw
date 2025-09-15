package handlers

import (
	"bytes"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sumanchapai/gw/scripts"
)

type ActionRequest struct {
	Action    string `json:"action"`
	CommitMsg string `json:"commitmsg"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // allow all
}

type WebsocketWriter struct {
	b bytes.Buffer
	c *websocket.Conn
}

func (w WebsocketWriter) Write(p []byte) (n int, err error) {
	err = w.c.WriteMessage(websocket.TextMessage, p)
	return len(p), err
}

// GitAction handler streams the result of running the provided git-action
func GitAction(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	defer conn.Close()
	var req ActionRequest

	if err := conn.ReadJSON(&req); err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("Invalid JSON"+err.Error()))
		return
	}

	stdout := WebsocketWriter{bytes.Buffer{}, conn}
	stderr := WebsocketWriter{bytes.Buffer{}, conn}

	// Get the appropriate script name
	switch req.Action {
	case "commit":
		err := scripts.CommitAll(req.CommitMsg, &stdout, &stderr)
		if err != nil {
			log.Println("Error running CommitAll", "err:", err)
			conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			return
		}
	default:
		log.Println("Invalid action received", req.Action)
		conn.WriteMessage(websocket.TextMessage, []byte("Invalid action received"))
		return
	}
}
