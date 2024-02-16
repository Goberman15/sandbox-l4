package server

import "github.com/gofiber/fiber/v2"

func (s *Server) RegisterRouter() {
	api := s.app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Author": "Akbar Ramadhan",
		})
	})

	itemRouter := api.Group("/items")	
	itemDetailRouter := api.Group("/item-details")	

	registerItemRouter(itemRouter, s.db)
	registerItemDetailRouter(itemDetailRouter, s.db)
}
