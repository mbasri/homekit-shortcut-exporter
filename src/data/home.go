package data

import "fmt"

type HomeManager interface {
	NewHome(name string) (*Home, error)
	GetHome(id int) (*Home, error)

	UpdateName(name string) error
	Delete() error
}

type Home struct {
	Id         int
	Name       string
	CreatedAt  string
	LastUpdate string
}

func NewHome(name string) (*Home, error) {
	h := &Home{
		Name:      name,
		CreatedAt: "createdAt",
	}

	return h, nil
}

func (h *Home) UpdateName(name string) error {
	h.Name = name
	return nil
}

func GetHome(id int) (*Home, error) {
	return &Home{}, nil
}

func (h *Home) Delete() error {
	return nil
}

func (u Home) String() string {
	return fmt.Sprintf("Home: %d, Name: %s, CreatedAt: %s, LastUpdate: %s", u.Id, u.Name, u.CreatedAt, u.LastUpdate)
}
