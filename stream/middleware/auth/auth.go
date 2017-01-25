package auth

import (
	"fmt"
	"net/http"

	"github.com/sctlee/gipfeli/auth/api"
	"golang.org/x/net/context"
	"golang.org/x/net/websocket"
	"google.golang.org/grpc"
)

const (
	GipfeliAccessToken = "access_token"
	GipfeliAuthURL     = ":9010"
)

// Authorize ...
func Authorize(config *websocket.Config, req *http.Request) (err error) {
	err = req.ParseForm()
	if err != nil {
		return err
	}
	accessToken := req.Form.Get(GipfeliAccessToken)
	if accessToken == "" {
		return fmt.Errorf("require access_token")
	}

	conn, err := grpc.Dial(GipfeliAuthURL, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := api.NewGipfeliAuthServiceClient(conn)
	response, err := client.GetUserInfo(context.Background(), &api.GetUserInfoRequest{Authorization: accessToken})
	if err != nil {
		return err
	}
	req.Form.Set("user_id", response.UserId)

	return nil
}

// GetUserID ...
func GetUserID(req *http.Request) string {
	err := req.ParseForm()
	if err != nil {
		return ""
	}
	return req.Form.Get("user_id")
}
