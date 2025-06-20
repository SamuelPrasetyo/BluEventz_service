package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	eventsRoutes(api)

	// api.Get("/users", controllers.GetUsers)
	// api.Post("/users", controllers.CreateUser)
}
