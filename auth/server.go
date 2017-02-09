package auth

import (
	"github.com/sctlee/gipfeli/auth/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Server ...
type Server struct{}

// GetUserInfo ...
func (s *Server) GetUserInfo(ctx context.Context, in *api.GetUserInfoRequest) (*api.GetUserInfoResponse, error) {
	if in.Authorization == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "argument error: require access_token")
	}
	return &api.GetUserInfoResponse{UserId: "123456", UserName: "sctlee"}, nil
}

// NewAuthServer ...
func NewAuthServer() *Server {
	return &Server{}
}
