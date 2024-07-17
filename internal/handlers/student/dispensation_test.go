package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestStudentHandler_CreateDispensation(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Post("/student/dispensation", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Post("/student/dispensation", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_GetDispensationByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/dispensation/:dispensationID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/dispensation/:dispensationID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_GetMyDispensations(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/dispensations", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/dispensations", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
