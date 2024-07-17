package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestStudentHandler_GetTask(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/task/:taskID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/task/:taskID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_SubmitTaskAssignment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Post("/student/task/:taskID/submit", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Post("/student/task/:taskID/submit", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_GetAssignment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/assignment/:assignmentID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/assignment/:assignmentID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_UpdateTaskSubmission(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Put("/student/task/:taskID/submit", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Put("/student/task/:taskID/submit", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
