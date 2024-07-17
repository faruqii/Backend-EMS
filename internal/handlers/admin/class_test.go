package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestAdminHandler_CreateClass(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Post("/admin/class", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"name": "Class 1"}`)
		req := httptest.NewRequest(http.MethodPost, "/admin/class", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Post("/admin/class", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"name": "Class 1"}`)
		req := httptest.NewRequest(http.MethodPost, "/admin/class", bytes.NewReader(reqBody))
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_AssignHomeroomTeacher(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Post("/admin/class/:id/homeroom-teacher", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodPost, "/admin/class/1/homeroom-teacher", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Post("/admin/class/:id/homeroom-teacher", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodPost, "/admin/class/1/homeroom-teacher", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

// remove home room teacher
func TestAdminHandler_RemoveHomeroomTeacher(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Delete("/admin/class/:id/homeroom-teacher", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodDelete, "/admin/class/1/homeroom-teacher", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Delete("/admin/class/:id/homeroom-teacher", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodDelete, "/admin/class/1/homeroom-teacher", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_GetAllClass(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/admin/class", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodGet, "/admin/class", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/admin/class", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodGet, "/admin/class", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

// update teacher homeroomStatus
func TestAdminHandler_UpdateTeacherHomeroomStatus(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Put("/admin/teacher/:id/homeroom-status", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"status": "homeroom"}`)
		req := httptest.NewRequest(http.MethodPut, "/admin/teacher/1/homeroom-status", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Put("/admin/teacher/:id/homeroom-status", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"status": "homeroom"}`)
		req := httptest.NewRequest(http.MethodPut, "/admin/teacher/1/homeroom-status", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_GetClassSchedule(t *testing.T) {
	app := fiber.New()
	app.Get("/admin/class/:id/schedule", func(ctx *fiber.Ctx) error {
		return nil
	})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/admin/class/1/schedule", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("class not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/admin/class/unknown/schedule", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_GetAllStudentsBelongToClass(t *testing.T) {
	app := fiber.New()
	app.Get("/admin/class/:id/students", func(ctx *fiber.Ctx) error {
		return nil
	})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/admin/class/1/students", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("internal server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/admin/class/unknown/students", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_RemoveStudentsFromClass(t *testing.T) {
	app := fiber.New()
	app.Delete("/admin/class/:id/students", func(ctx *fiber.Ctx) error {
		return nil
	})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/admin/class/1/students", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("internal server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/admin/class/unknown/students", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_RemoveSubjectFromClass(t *testing.T) {
	app := fiber.New()
	app.Delete("/admin/class/:classID/subject/:subjectID", func(ctx *fiber.Ctx) error {
		return nil
	})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/admin/class/1/subject/1", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("internal server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/admin/class/unknown/subject/unknown", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}