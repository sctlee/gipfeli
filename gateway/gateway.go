package gateway

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

type handshake func(config *websocket.Config, req *http.Request) (err error)

type Client struct {
	ws *websocket.Conn
}

func (self *Client) Read() (string, error) {
	var data string
	fmt.Println(self.ws)
	err := websocket.Message.Receive(self.ws, &data)
	return data, err
}

func (self *Client) Write(data string) {
	err := websocket.Message.Send(self.ws, data) // it has already included a lock
	if err != nil {
		fmt.Println(err)
		return
	}
}

type Gateway struct {
	// Clients map[string]*Client
	pending      chan *Client
	quiting      chan *Client // string : cid
	eventHandler EventHandler
	middlewares  []handshake
}

func NewGateway(handler EventHandler) *Gateway {
	return &Gateway{
		pending:      make(chan *Client),
		quiting:      make(chan *Client),
		eventHandler: handler,
		middlewares:  make([]handshake, 3),
	}
}

func (self *Gateway) Run(port int) {
	go self.listen()
	address := fmt.Sprintf("0.0.0.0:%d", port)

	http.Handle("/conn", self.genServer(self.accept))
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func (self *Gateway) RegisterMiddleware(hs handshake) {
	self.middlewares = append(self.middlewares, hs)
}

func (self *Gateway) genServer(h websocket.Handler) websocket.Server {
	hs := func(config *websocket.Config, req *http.Request) (err error) {
		for _, middleware := range self.middlewares {
			err := middleware(config, req)
			if err != nil {
				return err
			}
		}
		return
	}
	s := websocket.Server{Handler: h, Handshake: hs}
	return s
}

func (self *Gateway) accept(ws *websocket.Conn) {
	fmt.Println(ws)
	client := &Client{ws}

	self.pending <- client
	for {
		msg, err := client.Read()
		if err != nil {
			fmt.Println(err)
			self.quiting <- client
			break
		}
		self.eventHandler.Message(client, msg)
	}
}

func (self *Gateway) listen() {
	fmt.Println("Start listening...")
	for {
		fmt.Println("listen")
		select {
		case client := <-self.pending:
			self.eventHandler.Connect(client)
		case client := <-self.quiting:
			self.eventHandler.Disconnect(client)
		}
	}
}

func Start(eventHandler EventHandler, port int) {
	gateway := NewGateway(eventHandler)
	gateway.RegisterMiddleware(hs)
	gateway.Run(port)
}
