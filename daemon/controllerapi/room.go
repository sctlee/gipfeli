package controllerapi

import (
	"github.com/sctlee/gipfeli/daemon/api"
	"github.com/sctlee/gipfeli/daemon/store"
	"golang.org/x/net/context"
)

// ListRooms ...
func (s *Server) ListRooms(ctx context.Context, request *api.ListRoomsRequest) (*api.ListRoomsResponse, error) {
	var rooms []*api.Room

	rooms, err := store.ListRooms(s.store)
	if err != nil {
		return nil, err
	}

	return &api.ListRoomsResponse{
		Room: rooms,
	}, nil
}

// CreateRoom ...
func (s *Server) CreateRoom(ctx context.Context, request *api.CreateRoomRequest) (*api.CreateRoomResponse, error) {
	return &api.CreateRoomResponse{}, nil
}

// UpdateRoom ...
func (s *Server) UpdateRoom(context.Context, *api.UpdateRoomRequest) (*api.UpdateRoomResponse, error) {
	return &api.UpdateRoomResponse{}, nil
}

// DeleteRoom ...
func (s *Server) DeleteRoom(context.Context, *api.DeleteRoomRequest) (*api.DeleteRoomResponse, error) {
	return &api.DeleteRoomResponse{}, nil
}
