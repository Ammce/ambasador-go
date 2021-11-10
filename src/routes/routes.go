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

	// Orders
	adminAuthenticated.Get("/orders", controllers.Orders)

	// Links
	adminAuthenticated.Get("/user/:id/links", controllers.GetUserLinks)

	productAdmin := adminAuthenticated.Group("product")

	productAdmin.Post("/", controllers.CreateProduct)
	productAdmin.Patch("/:id", controllers.UpdateProduct)
	productAdmin.Get("/", controllers.Products)
	productAdmin.Get("/:id", controllers.Product)
	productAdmin.Delete("/:id", controllers.DeleteProduct)

	ambasador := api.Group("ambasador")
	ambasador.Post("/register", controllers.Register)
	ambasador.Post("/login", controllers.Login)

	ambasadorAuthenticated := ambasador.Use(middlewares.IsAuth)
	ambasadorAuthenticated.Get("/user", controllers.User)
	ambasadorAuthenticated.Post("/logout", controllers.Logout)
	ambasadorAuthenticated.Patch("/user", controllers.UpdateUser)

}
