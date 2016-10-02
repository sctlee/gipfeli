package gateway

import (
	"fmt"
	"net/http"

	auth_gw "github.com/sctlee/gipfeli/auth/api"
	daemon_gw "github.com/sctlee/gipfeli/daemon/api"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Start ...
func Start(port int) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := auth_gw.RegisterGipfeliAuthServiceHandlerFromEndpoint(ctx, mux, "0.0.0.0:9030", opts)
	if err != nil {
		fmt.Println(err)
	}
	err = daemon_gw.RegisterGipfeliRoomServiceHandlerFromEndpoint(ctx, mux, "0.0.0.0:9020", opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Start listening on 0.0.0.0:%d", port)
	address := fmt.Sprintf(":%d", port)
	http.ListenAndServe(address, mux)

}
