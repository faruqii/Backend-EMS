package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestTeacherHandler_GetAllAchivement(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/achivement/all", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/achivement/all", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestTeacherHandler_GetAchivementByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/achivement/:id", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/teacher/achivement/:id", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestTeacherHandler_UpdateAchievement(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Put("/teacher/achivement/:id/update", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Put("/teacher/achivement/:id/update", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

func TestTeacherHandler_DeleteAchivement(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Delete("/teacher/achivement/:id/delete", func(ctx *fiber.Ctx) error {
			return nil
		})
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Delete("/teacher/achivement/:id/delete", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}
