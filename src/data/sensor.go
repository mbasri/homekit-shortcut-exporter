package data

import "fmt"

type SensorManager interface {
	NewSensor(name string) (*Sensor, error)
	GetSensor(id int) (*Sensor, error)

	UpdateName(name string) error
	Delete() error
}

type Sensor struct {
	Id         int
	Name       string
	CreatedAt  string
	LastUpdate string
}

func NewSensor(name string) (*Sensor, error) {
	s := &Sensor{
		Name:      name,
		CreatedAt: "createdAt",
	}

	return s, nil
}

func (s *Sensor) UpdateName(name string) error {
	s.Name = name
	return nil
}

func Get(id int) (*Sensor, error) {
	return &Sensor{}, nil
}

func (s *Sensor) Delete() error {
	return nil
}

func (s Sensor) String() string {
	return fmt.Sprintf("Sensor: %d, Name: %s, CreatedAt: %s, LastUpdate: %s", s.Id, s.Name, s.CreatedAt, s.LastUpdate)
}
