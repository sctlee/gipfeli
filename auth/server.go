package auth

import (
	"github.com/sctlee/gipfeli/auth/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// server is used to implement helloworld.GreeterServer.
type Server struct{}

// SayHello implements helloworld.GreeterServer
func (s *Server) GetUserInfo(ctx context.Context, in *api.GetUserInfoRequest) (*api.GetUserInfoResponse, error) {
	if in.Authorization == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "argument error: require access_token")
	}
	return &api.GetUserInfoResponse{UserId: "123456", UserName: "sctlee"}, nil
}

func NewAuthServer() *Server {
	return &Server{}
}
