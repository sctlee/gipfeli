package websocket

import (
	"fmt"
	ws "golang.org/x/net/websocket"
)

// Client ...
type Client struct {
	conn *ws.Conn
}

func (c *Client) Read() (string, error) {
	var data string
	err := ws.Message.Receive(c.conn, &data)
	return data, err
}

func (c *Client) Write(data string) {
	err := ws.Message.Send(c.conn, data) // it has already included a lock
	if err != nil {
		fmt.Println(err)
		return
	}
}
