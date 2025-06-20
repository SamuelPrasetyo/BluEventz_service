package routes

import (
	"github.com/gofiber/fiber/v2"
)

func eventsRoutes(group fiber.Router) {
	event := group.Group("/events")

	event.Get("/", events.getEvents)
}
