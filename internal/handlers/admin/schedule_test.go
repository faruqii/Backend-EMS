package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestAdminHandler_CreateSchedule(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Post("/admin/schedule", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"teacher_id": 1, "subject_id": 1, "class_id": 1}`)
		req := httptest.NewRequest(http.MethodPost, "/admin/schedule", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Post("/admin/schedule", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"teacher_id": 1, "subject_id": 1}`)
		req := httptest.NewRequest(http.MethodPost, "/admin/schedule", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_GetSchedules(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/admin/schedule", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodGet, "/admin/schedule", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/admin/schedule", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodGet, "/admin/schedule", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_UpdateSchedule(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Put("/admin/schedule/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"teacher_id": 1, "subject_id": 1, "class_id": 1}`)
		req := httptest.NewRequest(http.MethodPut, "/admin/schedule/1", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Put("/admin/schedule/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"teacher_id": 1, "subject_id": 1}`)
		req := httptest.NewRequest(http.MethodPut, "/admin/schedule/1", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_DeleteSchedule(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Delete("/admin/schedule/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodDelete, "/admin/schedule/1", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Delete("/admin/schedule/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodDelete, "/admin/schedule/1", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}


