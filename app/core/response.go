package core

import "github.com/gofiber/fiber/v2"

func JSONResponse(c *fiber.Ctx, status bool, message string, data interface{}) error {
	return c.JSON(fiber.Map{
		"responseStatus":  status,
		"responseMessage": message,
		"responseBody":    data,
	})
}
