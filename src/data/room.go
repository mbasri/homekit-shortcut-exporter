package data

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type RoomManager interface {
	UpdateName(name string) error
	Delete() error
}

type Room struct {
	Id         string
	Name       string
	CreatedAt  string
	LastUpdate string
}

func NewRoom(name string) (*Room, error) {
	r := &Room{
		Id:        uuid.NewString(),
		Name:      name,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05.000 -0700"),
	}

	return r, nil
}

func (r *Room) UpdateName(name string) error {
	r.Name = name
	r.LastUpdate = time.Now().Format("2006-01-02 15:04:05.000 -0700")
	return nil
}

func GetRoom(id string) (*Room, error) {
	return &Room{}, nil
}

func (r *Room) Delete() error {
	return nil
}

func (r Room) String() string {
	return fmt.Sprintf("Room: %d, Name: %s, CreatedAt: %s, LastUpdate: %s", r.Id, r.Name, r.CreatedAt, r.LastUpdate)
}
