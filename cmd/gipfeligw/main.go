package main

import (
	"github.com/sctlee/gipfeli/gateway"
	"github.com/sctlee/gipfeli/gateway/websocket"
)

func main() {
	go websocket.Start(9010)
	gateway.Start(9000)
}
