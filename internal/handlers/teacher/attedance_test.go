package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestTeacherHandler_CreateAttendance(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Post("/teacher/attendance/:subjectID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Post("/teacher/attendance/:subjectID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

}

func TestTeacherHandler_GetAttendanceBySubjectID(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/attendance/:subjectID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/attendance/:subjectID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestTeacherHandler_GetAttendanceByClassID(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/attendance/:classID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/attendance/:classID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestTeacherHandler_UpdateAttendance(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Put("/teacher/attendance/:id/update", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Put("/teacher/attendance/:id/update", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
