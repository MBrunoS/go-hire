package server

import (
	"github.com/mbrunos/go-hire/pkg/router"
)

type Server struct {
	Router router.Router
	Port   string
}

func NewServer(port string) *Server {
	return &Server{
		Router: router.NewDefaultRouter(),
		Port:   port,
	}
}

func (s *Server) Start() error {
	if err := s.Router.Serve(s.Port); err != nil {
		return err
	}

	return nil
}
