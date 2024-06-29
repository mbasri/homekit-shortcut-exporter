package main

import (
	"nmc-account/server"
	"nmc-account/store"
	"os"
)

func main() {

	if err := run(); err != nil {
		os.Exit(1)
	}
}

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
