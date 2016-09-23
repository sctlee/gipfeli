package daemon

import (
	"github.com/sctlee/gipfeli/daemon/api"
	"golang.org/x/net/context"
)

// Server ...
type Server struct{}

func (s *Server) ListRooms(context.Context, *api.ListRoomsRequest) (*api.ListRoomsResponse, error) {

}

func (s *Server) JoinRooms(context.Context, *api.JoinRoomsRequest) (*api.JoinRoomsResponse, error) {

}

func (s *Server) ExitRooms(context.Context, *api.ExitRoomsRequest) (*api.ExitRoomsResponse, error) {

}
