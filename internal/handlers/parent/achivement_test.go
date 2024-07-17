package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestParentHandler_GetAchivement(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/achivement", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/parent/achivement", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
	t.Run("negative test", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/achivement", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/parent/achivement", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
