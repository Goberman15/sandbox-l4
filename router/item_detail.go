package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/goberman15/sandbox-l4/controller"
	"github.com/goberman15/sandbox-l4/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func NewItemDetailRouter(db *sqlx.DB) *chi.Mux {
	c := controller.NewItemDetailController(repository.NewItemDetailRepo(db), repository.NewItemRepo(db))
	r := chi.NewRouter()

	r.Get("/{itemId}", c.ListItemDetailByItemId)
	r.Post("/", c.CreateItemDetail)
	r.Put("/{id}", c.UpdateItemDetail)
	r.Delete("/{id}", c.DeleteItemDetail)

	return r
}

func registerItemDetailRouter(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"Author": "Akbar Ramadhan",
			"Age": 27,
		})
	})
}


