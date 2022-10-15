package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		resp := struct {
			Message string
		}{
			Message: "Hello World",
		}

		return c.JSON(resp)
	})

	app.Listen(":8080")
}
