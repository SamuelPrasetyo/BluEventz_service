package routes

import (
	"github.com/SamuelPrasetyo/BluEventz_service/app/modules/events/controller"
	"github.com/gofiber/fiber/v2"
)

func eventsRoutes(group fiber.Router) {
	// Group
	event := group.Group("/events")

	// Endpoint
	event.Get("/get-events", controller.GetEvents)
}
