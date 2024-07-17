package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestTeacherHandler_CreateViolation(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Post("/teacher/violation", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Post("/teacher/violation", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestTeacherHandler_GetAllViolation(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/violations", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/violations", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

}

func TestTeacherHandler_GetViolationByID(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/violation/:violationID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/violation/:violationID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestTeacherHandler_GetViolationByStudentID(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/violation/student/:studentID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/violation/student/:studentID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
