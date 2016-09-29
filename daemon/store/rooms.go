package store

import (
	"encoding/json"
	"fmt"

	"github.com/sctlee/gipfeli/daemon/api"
)

var tableRoom = "room"

func init() {
	registry(&ObjectConfig{
		Name: tableRoom,
		Table: &TableSchema{
			Name:    tableRoom,
			Indexes: map[string]*IndexSchema{},
		},
		Objecter: func() Object {
			return roomEntry{}
		},
	})
}

type roomEntry struct {
	*api.Room
}

func (r roomEntry) ID() string {
	return r.Room.Id
}

func (r roomEntry) JSONData() (string, error) {
	jsonData, err := json.Marshal(r.Room)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (r roomEntry) SetContent(jsonData string) error {
	err := json.Unmarshal([]byte(jsonData), r.Room)
	fmt.Println("ttttt", r.Room)
	return err
}

// ListRooms ...
func ListRooms(tx ReadTx, bys ...By) ([]*api.Room, error) {
	roomList := []*api.Room{}
	appendResult := func(o Object) {
		roomList = append(roomList, o.(roomEntry).Room)
	}

	err := tx.find(tableRoom, appendResult, bys...)
	return roomList, err
}

// CreateRoom ...
func CreateRoom() {

}

// UpdateRoom ...
func UpdateRoom() {

}

// DeleteRoom ...
func DeleteRoom() {

}
