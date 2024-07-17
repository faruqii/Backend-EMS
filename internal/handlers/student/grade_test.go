package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestStudentHandler_GetGradeByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/grade/:id", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/grade/:id", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_GetMyGrades(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/grades", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/grades", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
