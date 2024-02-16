package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	app *fiber.App
	db  *sqlx.DB
}

func NewServer(db *sqlx.DB) *Server {
	app := fiber.New()
	return &Server{
		app: app,
		db:  db,
	}
}

func (s *Server) Start() error {
	return s.app.Listen(":3000")
}
