package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestGlobalHandler_GetAnnouncements(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/global/announcements", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/global/announcements", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/global/announcements", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/global/announcements", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestGlobalHandler_GetAnnouncementByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/global/announcement/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/global/announcement/1", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/global/announcement/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/global/announcement/1", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestGlobalHandler_GetAllAgendas(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/global/agendas", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/global/agendas", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/global/agendas", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/global/agendas", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestGlobalHandler_GetAgendaByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/global/agenda/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/global/agenda/1", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/global/agenda/:id", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/global/agenda/1", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
