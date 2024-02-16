package server

import (
	"github.com/goberman15/sandbox-l4/controller"
	"github.com/goberman15/sandbox-l4/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func registerItemDetailRouter(r fiber.Router, db *sqlx.DB) {
	c := controller.NewItemDetailController(repository.NewItemDetailRepo(db), repository.NewItemRepo(db))

	r.Post("/", c.CreateItemDetail)
	r.Get("/:itemId", c.ListItemDetailByItemId)
	r.Put("/:id", c.UpdateItemDetail)
	r.Delete("/:id", c.DeleteItemDetail)
}
