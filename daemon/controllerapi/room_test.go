package controllerapi

import (
	"fmt"
	"github.com/sctlee/gipfeli/daemon/api"
	"golang.org/x/net/context"
	"testing"
)

func TestListServices(t *testing.T) {
	ts := newTestServer(t)
	r, err := ts.Client.ListRooms(context.Background(), &api.ListRoomsRequest{})
	t.Error(err)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("asdf", r)
}
