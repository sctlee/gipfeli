package gipfeli

import (
	"net"
)

type Service struct {
	Label string
	conn  net.Conn
}
