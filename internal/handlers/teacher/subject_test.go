package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestTeacherHandler_CountStudent(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/subject/:subjectID/student/count", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/subject/:subjectID/student/count", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestTeacherHandler_GetMySubjects(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/subjects", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/subjects", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestTeacherHandler_CreateSubjectMatter(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Post("/teacher/subject/:subjectID/matter", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Post("/teacher/subject/:subjectID/matter", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestTeacherHandler_GetSubjectMatterBySubjectID(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/subject/:subjectID/matters", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/subject/:subjectID/matters", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestTeacherHandler_GetDetailSubjectMatter(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/subject/matter/:subjectMatterID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/subject/matter/:subjectMatterID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestTeacherHandler_UpdateSubjectMatter(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Put("/teacher/subject/matter/:subjectMatterID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Put("/teacher/subject/matter/:subjectMatterID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestTeacherHandler_DeleteSubjectMatter(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		app := fiber.New()
		app.Delete("/teacher/subject/matter/:subjectMatterID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("Negative Test", func(t *testing.T) {
		app := fiber.New()
		app.Delete("/teacher/subject/matter/:subjectMatterID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
