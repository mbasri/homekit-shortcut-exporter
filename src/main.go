package main

import (
	"fmt"
	"homekit-shortcut-exporter/data"
)

func main() {

	var h data.HomeManager
	var r data.RoomManager
	var s data.SensorManager

	h, _ = data.NewHome("home")
	r, _ = data.NewRoom("room")
	s, _ = data.NewSensor("sensor")

	fmt.Printf("%s", h)
	fmt.Printf("%s", r)
	fmt.Printf("%s", s)

	/*if err := run(); err != nil {
		os.Exit(1)
	}*/
}

/*
func run() error {
	srv := server.New()
	srv.Store = &store.DBStore{}

	err := srv.Store.Open()
	if err != nil {
		return err
	}
	defer srv.Store.Close()

	return nil
}
*/
