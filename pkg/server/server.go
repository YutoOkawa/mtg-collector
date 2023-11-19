package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app  *fiber.App
	port string
}

func NewServer(port int) *Server {
	app := fiber.New()
	return &Server{
		app:  app,
		port: fmt.Sprintf(":%d", port),
	}
}

func (s *Server) Start(errCh chan<- error) {
	if err := s.app.Listen(s.port); err != nil {
		errCh <- err
	}
}
