package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestTeacherHandler_GetMyProfile(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/profile", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/profile", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
