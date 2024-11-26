package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello and uhmmmm, just create cool thing just creat enew thing")
	})

	app.Listen(":3000")
}
