package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *StudentHandler) MyProfile(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	userID := ctx.Locals("user").(string)

	profile, err := h.studentService.MyProfile(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get profile",
		"data":    profile,
	})
}
