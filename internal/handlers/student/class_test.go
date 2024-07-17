package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestStudentHandler_GetClass(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/class", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/class", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_GetSubjects(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/subjects", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/subjects", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
