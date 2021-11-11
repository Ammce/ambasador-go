package main

import (
	"github.com/Ammce/ambasador-go/src/database"
	"github.com/Ammce/ambasador-go/src/middlewares"
	"github.com/Ammce/ambasador-go/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()
	database.AutoMigrate()
	database.SetupRedis()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	middlewares.Logger(app)

	routes.Setup(app)

	app.Listen(":8080")
}
