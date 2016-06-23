package gateway

import (
	"fmt"
	"net/http"

	uuid "github.com/nu7hatch/gouuid"
	"golang.org/x/net/websocket"
)

type Client struct {
	Id string
	ws *websocket.Conn
}

func NewClient(ws *websocket.Conn) (*Client, error) {
	cid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	client := &Client{
		Id: cid.String(),
		ws: ws,
	}
	return client, nil
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
	Pending      chan *Client
	Quiting      chan *Client // string : cid
	eventHandler EventHandler
}

func NewGateway(handler EventHandler) *Gateway {
	return &Gateway{
		Pending:      make(chan *Client),
		Quiting:      make(chan *Client),
		eventHandler: handler,
	}
}

func (self *Gateway) Run(port int) {
	go self.listen()

	http.Handle("/conn", websocket.Handler(self.accept))
	address := fmt.Sprintf("0.0.0.0:%d", port)
	fmt.Println(address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func (self *Gateway) accept(ws *websocket.Conn) {
	fmt.Println(ws)
	client, err := NewClient(ws)
	if err != nil {
		fmt.Println(err)
		return
	}

	self.Pending <- client
	for {
		msg, err := client.Read()
		if err != nil {
			fmt.Println(err)
			self.Quiting <- client
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
		case client := <-self.Pending:
			self.eventHandler.Connect(client)
		case client := <-self.Quiting:
			self.eventHandler.Disconnect(client)
		}
	}
}

func Start(eventHandler EventHandler, port int) {
	gateway := NewGateway(eventHandler)
	gateway.Run(port)
}
