/* Tempat semua endpoint HTTP/REST berada */

package events

import (
	"github.com/SamuelPrasetyo/BluEventz_service/utils"
	"github.com/gofiber/fiber/v2"
)

func GetEventsHandler(c *fiber.Ctx) error {
	data, err := GetEventsService()

	// Handle null atau zero value
	if err != nil {
		return utils.ServerError(c, "Gagal mengambil data Events")
	}

	// Kirim response success
	return utils.Success(c, "Berhasil mengambil data event", data)
}
