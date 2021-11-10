package routes

import (
	"github.com/Ammce/ambasador-go/src/controllers"
	"github.com/Ammce/ambasador-go/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")

	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)

	adminAuthenticated := admin.Use(middlewares.IsAuth)

	adminAuthenticated.Get("/user", controllers.User)
	adminAuthenticated.Post("/logout", controllers.Logout)
	adminAuthenticated.Patch("/user", controllers.UpdateUser)
	adminAuthenticated.Get("/ambasadors", controllers.GetAmbasadors)

}
