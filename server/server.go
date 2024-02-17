package server

import (
	"errors"

	"github.com/goberman15/sandbox-l4/customErr"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	app *fiber.App
	db  *sqlx.DB
}

func NewServer(db *sqlx.DB) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
	app.Use(logger.New())
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"sandbox": "unbroken",
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return customErr.NewUnauthorizedError("you are not allowed to perform this action")
		},
	}))
	return &Server{
		app: app,
		db:  db,
	}
}

func (s *Server) Start() error {
	return s.app.Listen(":8080")
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	var e customErr.CustomError
	if errors.As(err, &e) {
		return ctx.Status(e.StatusCode()).JSON(fiber.Map{
			"message": e.Error(),
		})
	}

	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
}
