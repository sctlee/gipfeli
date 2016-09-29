package controllerapi

import (
	"github.com/sctlee/gipfeli/daemon/store"
)

// Server ...
type Server struct {
	store *store.RedisStore
}

// NewServer ...
func NewServer() *Server {
	return &Server{store.NewRedisStore()}
}
