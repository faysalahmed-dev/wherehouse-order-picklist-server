package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GlobalErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	var message string
	var stack interface{}
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
		stack = e.Error()
	}
	if err != nil {
		// In case the SendFile fails
		return ctx.Status(code).JSON(fiber.Map{
			"error":   true,
			"message": message,
			"stack":   stack,
		})
	}
	return nil
}
