package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestParentHandler_GetMyStudents(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/students", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/students", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestParentHandler_GetStudentDetail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/student/:studentID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/student/:studentID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
