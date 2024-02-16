package router

import "github.com/gofiber/fiber/v2"

func RegisterRouter(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Author": "Akbar Ramadhan",
		})
	})

	itemRouter := api.Group("/items")	
	itemDetailRouter := api.Group("/item-details")	

	registerItemRouter(itemRouter)
	registerItemDetailRouter(itemDetailRouter)
}
