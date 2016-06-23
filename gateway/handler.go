package gateway

type EventHandler interface {
	Connect(client *Client)
	Disconnect(client *Client)
	Message(client *Client, msg string)
}
