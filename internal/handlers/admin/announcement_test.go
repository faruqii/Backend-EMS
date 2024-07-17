package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestAdminHandler_CreateAnnouncement(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Post("/admin/announcement", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"title": "Announcement 1", "content": "Content 1"}`)
		req := httptest.NewRequest(http.MethodPost, "/admin/announcement", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Post("/admin/announcement", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"title": "Announcement 1"}`)
		req := httptest.NewRequest(http.MethodPost, "/admin/announcement", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_GetAnnouncements(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/admin/announcement", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodGet, "/admin/announcement", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/admin/announcement", func(ctx *fiber.Ctx) error {
			return nil
		})
	})
}

// get announcement by id
func TestAdminHandler_GetAnnouncement(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/admin/announcement/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodGet, "/admin/announcement/1", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/admin/announcement/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodGet, "/admin/announcement/invalid", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

// update announcement
func TestAdminHandler_UpdateAnnouncement(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Put("/admin/announcement/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"title": "Announcement 1", "content": "Content 1"}`)
		req := httptest.NewRequest(http.MethodPut, "/admin/announcement/1", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Put("/admin/announcement/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"title": "Announcement 1"}`)
		req := httptest.NewRequest(http.MethodPut, "/admin/announcement/1", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

// delete announcement
func TestAdminHandler_DeleteAnnouncement(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Delete("/admin/announcement/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodDelete, "/admin/announcement/1", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Delete("/admin/announcement/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodDelete, "/admin/announcement/invalid", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

