package server

import (
	"net/http"

	"github.com/mbrunos/go-hire/pkg/router"
)

type Server struct {
	Router   router.Router
	Handlers map[string]http.HandlerFunc
	Port     string
}

func NewServer(port string) *Server {
	return &Server{
		Router:   router.NewDefaultRouter(),
		Handlers: make(map[string]http.HandlerFunc),
		Port:     port,
	}
}

func (s *Server) AddRoute(route string, handler http.HandlerFunc) {
	s.Handlers[route] = handler
}

func (s *Server) AddRoutes(configure func(router.Router)) {
	configure(s.Router)
}

func (s *Server) Start() error {
	for route, handler := range s.Handlers {
		s.Router.Handle(route, handler)
	}

	if err := s.Router.Serve(s.Port); err != nil {
		return err
	}

	return nil
}
