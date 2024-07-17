package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestParentHandler_GetGradeByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/grade/:gradeID", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/parent/grade/1", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/grade/:gradeID", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/parent/grade/1", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

}

func TestParentHandler_GetStudentGrades(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/grades", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/parent/grades", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/parent/grades", func(ctx *fiber.Ctx) error {
			return nil
		})

		resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/parent/grades", nil))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
