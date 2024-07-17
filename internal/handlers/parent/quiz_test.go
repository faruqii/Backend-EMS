package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestParentHandler_GetQuizAssignment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/quiz/:quizID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/quiz/:quizID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
