package stream

import (
	"github.com/sctlee/gipfeli/stream/middleware/auth"
	"github.com/sctlee/gipfeli/stream/websocket"
)

// StartWithHandler ...
func StartWithHandler(eventHandler websocket.EventHandler, port int) {
	websocket := websocket.NewWebSocket(eventHandler)
	websocket.RegisterMiddleware(auth.Authorize)
	websocket.Run(port)
}

// Start ...
func Start(port int) {
	StartWithHandler(&websocket.PureEventHandler{}, port)
}
