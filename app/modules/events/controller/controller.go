package controller

import (
	"github.com/SamuelPrasetyo/BluEventz_service/app/core"
	"github.com/SamuelPrasetyo/BluEventz_service/app/modules/events/service"
	"github.com/gofiber/fiber/v2"
)

func GetEvents(c *fiber.Ctx) error {
	data, err := service.GetEventsService()

	// Handle null atau zero value
	if err != nil {
		return core.ServerError(c, "Gagal mengambil data Events")
	}

	// Kirim response success
	return core.Success(c, "Berhasil mengambil data event", data)
}
