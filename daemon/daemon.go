package daemon

import (
	"fmt"
	"log"
	"net"

	"github.com/sctlee/gipfeli/daemon/api"
	"github.com/sctlee/gipfeli/daemon/controllerapi"
	"google.golang.org/grpc"
)

// Start ...
func Start(port int) {
	fmt.Printf("Start listening on 0.0.0.0:%d", port)
	address := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterGipfeliRoomServiceServer(s, controllerapi.NewServer())
	s.Serve(lis)
}
