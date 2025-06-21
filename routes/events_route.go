package routes

import (
	"github.com/SamuelPrasetyo/BluEventz_service/internal/events"
	"github.com/gofiber/fiber/v2"
)

func eventsRoutes(group fiber.Router) {
	// Group
	event := group.Group("/events")

	// Endpoint
	event.Get("/get-events", events.GetEventsHandler)
}
