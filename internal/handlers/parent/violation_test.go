package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestParentHandler_GetStudentViolation(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/student/:studentID/violation", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/student/:studentID/violation", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestParentHandler_GetViolationByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/violation/:violationID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/violation/:violationID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
