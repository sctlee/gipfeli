package websocket

import (
	"github.com/sctlee/gipfeli/gateway/middleware/auth"
)

// StartWithHandler ...
func StartWithHandler(eventHandler EventHandler, port int) {
	websocket := NewWebSocket(eventHandler)
	websocket.RegisterMiddleware(auth.Authorize)
	websocket.Run(port)
}

// Start ...
func Start(port int) {
	StartWithHandler(&PureEventHandler{}, port)
}
