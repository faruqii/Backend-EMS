package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestTeacherHandler_GetTodaySchedule(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/schedule", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/schedule", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestTeacherHandler_GetAllTeacherSchedule(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/schedules", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/schedules", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
