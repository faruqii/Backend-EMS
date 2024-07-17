package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestStudentHandler_GetSubjectByClassID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/subject/:classID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/subject/:classID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_GetDetailSubject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/subject/detail/:subjectID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/subject/detail/:subjectID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_GetSubjectMatterBySubjectID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/subject/matter/:subjectID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/subject/matter/:subjectID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestStudentHandler_GetDetailSubjectMatter(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/subject/matter/detail/:matterID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/student/subject/matter/detail/:matterID", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
