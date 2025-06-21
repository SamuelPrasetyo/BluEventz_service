package utils

import (
	"github.com/gofiber/fiber/v2"
)

// Struct yang konsisten dipakai untuk semua response
type Response struct {
	ResponseStatus  bool        `json:"responseStatus"`
	ResponseMessage string      `json:"responseMessage"`
	ResponseBody    interface{} `json:"responseBody,omitempty"`
}

// Fungsi dasar (jika ingin tetap fleksibel)
func JSONResponse(c *fiber.Ctx, statusCode int, status bool, message string, data interface{}) error {
	return c.Status(statusCode).JSON(Response{
		ResponseStatus:  status,
		ResponseMessage: message,
		ResponseBody:    data,
	})
}

// Success response
func Success(c *fiber.Ctx, message string, data interface{}) error {
	return JSONResponse(c, fiber.StatusOK, true, message, data)
}

// Bad Request
func BadRequest(c *fiber.Ctx, message string) error {
	return JSONResponse(c, fiber.StatusBadRequest, false, message, nil)
}

// Unauthorized
func Unauthorized(c *fiber.Ctx, message string) error {
	return JSONResponse(c, fiber.StatusUnauthorized, false, message, nil)
}

// Internal Server Error
func ServerError(c *fiber.Ctx, message string) error {
	return JSONResponse(c, fiber.StatusInternalServerError, false, message, nil)
}
