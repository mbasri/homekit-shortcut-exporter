package data

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type SensorManager interface {
	UpdateName(name string) error
	Delete() error
}

type Sensor struct {
	Id         string
	Name       string
	CreatedAt  string
	LastUpdate string
}

func NewSensor(name string) (*Sensor, error) {
	s := &Sensor{
		Id:        uuid.NewString(),
		Name:      name,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05.000 -0700"),
	}

	return s, nil
}

func (s *Sensor) UpdateName(name string) error {
	s.Name = name
	s.LastUpdate = time.Now().Format("2006-01-02 15:04:05.000 -0700")
	return nil
}

func Get(id string) (*Sensor, error) {
	return &Sensor{}, nil
}

func (s *Sensor) Delete() error {
	return nil
}

func (s Sensor) String() string {
	return fmt.Sprintf("Sensor: %d, Name: %s, CreatedAt: %s, LastUpdate: %s", s.Id, s.Name, s.CreatedAt, s.LastUpdate)
}
