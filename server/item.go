package server

import (
	"github.com/goberman15/sandbox-l4/controller"
	"github.com/goberman15/sandbox-l4/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func registerItemRouter(r fiber.Router, db *sqlx.DB) {
	c := controller.NewItemController(repository.NewItemRepo(db))
	r.Get("/", c.ListItems)
	r.Get("/:id", c.GetItemById)
	r.Post("/", c.CreateItem)
	r.Patch("/:id/status", c.UpdateItemStatus)
	r.Patch("/:id/amount", c.UpdateItemAmount)
	r.Delete("/:id", c.DeleteItem)
}
