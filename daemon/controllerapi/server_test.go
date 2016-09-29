package controllerapi

import (
	"github.com/sctlee/gipfeli/daemon/api"
	"google.golang.org/grpc"
	// "io/ioutil"
	"net"
	"testing"
	"time"
)

type testServer struct {
	Server     *Server
	Client     api.GipfeliRoomServiceClient
	grpcServer *grpc.Server
	clientConn *grpc.ClientConn
}

func (ts *testServer) Stop() {
	_ = ts.clientConn.Close()
	ts.grpcServer.Stop()
}

func newTestServer(t *testing.T) *testServer {
	ts := &testServer{}

	ts.Server = NewServer()

	lis, _ := net.Listen("tcp", ":5200")

	ts.grpcServer = grpc.NewServer()
	api.RegisterGipfeliRoomServiceServer(ts.grpcServer, ts.Server)
	go func() {
		// Serve will always return an error (even when properly stopped).
		// Explicitly ignore it.
		_ = ts.grpcServer.Serve(lis)
	}()

	time.Sleep(5 * time.Second)

	conn, _ := grpc.Dial(":5200", grpc.WithInsecure())

	ts.Client = api.NewGipfeliRoomServiceClient(conn)

	return ts
}
