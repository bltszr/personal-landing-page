package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/static", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("./templates/index.html", fiber.Map{})
	})

	app.Listen(":3000")
}
