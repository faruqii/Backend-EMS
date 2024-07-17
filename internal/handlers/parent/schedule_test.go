package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestParentHandler_GetSchedule(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/schedule", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/schedule", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
