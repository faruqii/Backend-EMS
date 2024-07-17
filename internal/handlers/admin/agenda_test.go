package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestAdminHandler_CreateAgenda(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Post("/admin/agenda", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"title": "Agenda 1", "description": "Description 1", "date": "2021-08-01"}`)
		resp, err := app.Test(httptest.NewRequest(http.MethodPost, "/admin/agenda", bytes.NewReader(reqBody)))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Post("/admin/agenda", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"title": "Agenda 1", "description": "Description 1"}`)
		resp, err := app.Test(httptest.NewRequest(http.MethodPost, "/admin/agenda", bytes.NewReader(reqBody)))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_GetAgendas(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/admin/agendas", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/admin/agendas", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/admin/agendas", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/admin/agendas", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_GetAgenda(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/admin/agenda/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/admin/agenda/1", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/admin/agenda/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/admin/agenda/1", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_UpdateAgenda(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Put("/admin/agenda/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"title": "Agenda 1", "description": "Description 1", "date": "2021-08-01"}`)
		resp, err := app.Test(httptest.NewRequest(http.MethodPut, "/admin/agenda/1", bytes.NewReader(reqBody)))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Put("/admin/agenda/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"title": "Agenda 1", "description": "Description 1"}`)
		resp, err := app.Test(httptest.NewRequest(http.MethodPut, "/admin/agenda/1", bytes.NewReader(reqBody)))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_DeleteAgenda(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Delete("/admin/agenda/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodDelete, "/admin/agenda/1", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Delete("/admin/agenda/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodDelete, "/admin/agenda/1", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

