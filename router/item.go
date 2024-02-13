package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/goberman15/sandbox-l4/controller"
	"github.com/goberman15/sandbox-l4/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func NewItemRouter(db *sqlx.DB) *chi.Mux {

	c := controller.NewItemController(repository.NewItemRepo(db))

	r := chi.NewRouter()

	r.Get("/", c.ListItems)
	r.Get("/{id}", c.GetItemById)
	r.Post("/", c.CreateItem)
	r.Patch("/{id}/status", c.UpdateItemStatus)
	r.Patch("/{id}/amount", c.UpdateItemAmount)
	r.Delete("/{id}", c.DeleteItem)

	return r
}

func registerItemRouter(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"Author": "Akbar Ramadhan",
			"Age":    27,
		})
	})
}
