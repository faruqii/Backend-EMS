package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestStudentHandler_MyProfile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/profile", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/profile", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
