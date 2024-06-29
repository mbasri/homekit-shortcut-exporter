package server

import "nmc-account/store"

type Server struct {
	Store store.Store
}

func New() *Server {
	return &Server{}
}
