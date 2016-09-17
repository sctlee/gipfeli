package websocket

import (
	"log"

	"github.com/sctlee/gipfeli/gateway/middleware/auth"
	// "github.com/gorilla/mux"
)

// PureEventHandler ...
type PureEventHandler struct {
}

// Connect ..
func (peh *PureEventHandler) Connect(client *Client) {
	// vars := mux.Vars(client.conn.Request())
	// roomID := vars["room"]
	userID := auth.GetUserID(client.conn.Request())
	log.Printf("%s connect\n", userID)
}

// Disconnect ..
func (peh *PureEventHandler) Disconnect(client *Client) {
	log.Println("one client disconnect")
}

// Message ...
func (peh *PureEventHandler) Message(client *Client, msg string) {
	log.Println(msg)
}
