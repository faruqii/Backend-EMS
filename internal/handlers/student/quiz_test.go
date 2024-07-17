package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestStudentHandler_GetQuiz(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/quiz/:id", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/quiz/:id", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_GetQuizQuestions(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/quiz/:id/questions", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/quiz/:id/questions", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_SubmitQuizAnswer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Post("/student/quiz/:id/answer", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Post("/student/quiz/:id/answer", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_GetMyQuizGrade(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/quiz/:id/grade", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/quiz/:id/grade", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_GetMyQuizGrades(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/quiz/grades", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/quiz/grades", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_GetMyQuizSubmission(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/quiz/:id/submission", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/quiz/:id/submission", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
