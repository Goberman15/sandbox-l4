package router

import "github.com/gofiber/fiber/v2"

func RegisterRouter(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Author": "Akbar Ramadhan",
		})
	})

	itemRouter := r.Group("/items")	
	itemDetailRouter := r.Group("/item-details")	

	registerItemRouter(itemRouter)
	registerItemDetailRouter(itemDetailRouter)
}
