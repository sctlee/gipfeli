package auth

import (
	"log"
	"net"

	"github.com/sctlee/gipfeli/auth/api"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
func Start(port int) {
	lis, err := net.Listen("tcp", ":"+string(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterGipfeliAuthServiceServer(s, NewAuthServer())
	s.Serve(lis)
}
