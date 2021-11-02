package routes

import (
	"github.com/Ammce/ambasador-go/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/admin/register", controllers.Register)
}
