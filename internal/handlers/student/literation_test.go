package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestStudentHandler_InsertLiteration(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Post("/student/literation", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Post("/student/literation", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_GetLiterationByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/literation/:literationID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/literation/:literationID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_GetLiterationByStudentID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/literation/student/:studentID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/literation/student/:studentID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
