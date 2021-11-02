package main

import (
	"github.com/Ammce/ambasador-go/src/database"
	"github.com/Ammce/ambasador-go/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()
	database.AutoMigrate()

	app := fiber.New()
	routes.Setup(app)

	app.Listen(":8080")
}
