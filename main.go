package main

import (
	"github.com/Ammce/ambasador-go/src/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Helloo")
	})

	app.Listen(":8080")
}
