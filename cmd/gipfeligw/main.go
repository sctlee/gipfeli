package main

import (
	"fmt"
	"github.com/sctlee/gipfeli/gateway"
)

type EventHandler struct {
}

func (self *EventHandler) Connect(client *gateway.Client) {
	fmt.Println("one client connect")
}

func (self *EventHandler) Disconnect(client *gateway.Client) {
	fmt.Println("one client disconnect")
}

func (self *EventHandler) Message(client *gateway.Client, msg string) {
	fmt.Println(msg)
}

func main() {
	handler := &EventHandler{}
	gateway.Start(handler, 5000)
}
