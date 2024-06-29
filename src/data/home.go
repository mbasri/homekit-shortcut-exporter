package data

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type HomeManager interface {
	UpdateName(name string) error
	Delete() error
}

type Home struct {
	Id         string
	Name       string
	CreatedAt  string
	LastUpdate string
}

func NewHome(name string) (*Home, error) {
	h := &Home{
		Id:        uuid.NewString(),
		Name:      name,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05.000 -0700"),
	}

	return h, nil
}

func (h *Home) UpdateName(name string) error {
	h.Name = name
	h.LastUpdate = time.Now().Format("2006-01-02 15:04:05.000 -0700")
	return nil
}

func GetHome(id string) (*Home, error) {
	return &Home{}, nil
}

func (h *Home) Delete() error {
	return nil
}

func (h Home) String() string {
	return fmt.Sprintf("Home: %d, Name: %s, CreatedAt: %s, LastUpdate: %s", h.Id, h.Name, h.CreatedAt, h.LastUpdate)
}
