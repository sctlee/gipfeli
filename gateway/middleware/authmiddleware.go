package middleware

import (
	"fmt"
	"net/http"

	"github.com/sctlee/gipfeli/auth/api"
	"golang.org/x/net/context"
	"golang.org/x/net/websocket"
	"google.golang.org/grpc"
)

const (
	GIPFELI_TOKEN_HEADER = "X-GIPFELI-TOKEN"
	GIPFELI_USER_HEADER  = "X-GIPFELI-USER"
	GIPFELI_AUTH_URL     = "gipfeli_auth:5010"
)

func GetUserInfo(config *websocket.Config, req *http.Request) (err error) {
	access_token := req.Header.Get(GIPFELI_TOKEN_HEADER)
	if access_token == "" {
		return fmt.Errorf("require X-GIPFELI-TOKEN header")
	}

	conn, err := grpc.Dial(GIPFELI_AUTH_URL, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := api.NewGipfeliAuthServiceClient(conn)
	response, err := client.GetUserInfo(context.Background(), &api.GetUserInfoRequest{Authorization: access_token})
	if err != nil {
		return err
	}
	req.Header.Set(GIPFELI_USER_HEADER, response.UserId)

	return nil
}
