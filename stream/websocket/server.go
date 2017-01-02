package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	ws "golang.org/x/net/websocket"
)

type handshake func(config *ws.Config, req *http.Request) (err error)

// EventHandler ...
type EventHandler interface {
	Connect(client *Client)
	Disconnect(client *Client)
	Message(client *Client, msg string)
}

// WebSocket ...
type WebSocket struct {
	// Clients map[string]*Client
	pending      chan *Client
	quiting      chan *Client // string : cid
	eventHandler EventHandler
	middlewares  []handshake
}

// NewWebSocket ...
func NewWebSocket(handler EventHandler) *WebSocket {
	return &WebSocket{
		pending:      make(chan *Client),
		quiting:      make(chan *Client),
		eventHandler: handler,
		middlewares:  make([]handshake, 0),
	}
}

// Run ...
func (w *WebSocket) Run(port int) {
	go w.listen()

	r := mux.NewRouter()
	r.Handle("/conn/{room}", w.genServer(w.accept))

	address := fmt.Sprintf("0.0.0.0:%d", port)
	fmt.Println("Start listening...", address)
	err := http.ListenAndServe(address, r)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

// RegisterMiddleware ...
func (w *WebSocket) RegisterMiddleware(hs handshake) {
	w.middlewares = append(w.middlewares, hs)
}

func (w *WebSocket) genServer(h ws.Handler) ws.Server {
	hs := func(config *ws.Config, req *http.Request) (err error) {
		for _, middleware := range w.middlewares {
			err := middleware(config, req)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
		return
	}
	s := ws.Server{Handler: h, Handshake: hs}
	return s
}

func (w *WebSocket) accept(conn *ws.Conn) {
	client := &Client{conn}

	w.pending <- client
	for {
		msg, err := client.Read()
		if err != nil {
			fmt.Println(err)
			w.quiting <- client
			break
		}
		w.eventHandler.Message(client, msg)
	}
}

func (w *WebSocket) listen() {
	for {
		select {
		case client := <-w.pending:
			w.eventHandler.Connect(client)
		case client := <-w.quiting:
			w.eventHandler.Disconnect(client)
		}
	}
}
