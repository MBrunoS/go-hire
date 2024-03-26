package server

import (
	"github.com/mbrunos/go-hire/pkg/logger"
	"github.com/mbrunos/go-hire/pkg/router"
)

type Server struct {
	Router router.Router
	Port   string
	Logger logger.Logger
}

func NewServer(port string, logger logger.Logger) *Server {
	return &Server{
		Router: router.NewDefaultRouter(),
		Port:   port,
		Logger: logger,
	}
}

func (s *Server) Start() error {
	if err := s.Router.Serve(s.Port); err != nil {
		return err
	}

	return nil
}
