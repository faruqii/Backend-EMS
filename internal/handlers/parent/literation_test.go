package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestParentHandler_GetStudentLiterations(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/literations", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/literations", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestParentHandler_GetStudentLiterationDetail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/literation/:literationID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/literation/:literationID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
