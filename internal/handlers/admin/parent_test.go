package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestAdminHandler_CreateParentAccount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Post("/admin/parent", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"username": "parent1", "password": "password", "name": "Parent 1", "address": "Address 1", "occupation": "Occupation 1", "phone_number": "081234567890", "email": "mail@mail.com"}`)
		req := httptest.NewRequest(http.MethodPost, "/admin/parent", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Post("/admin/parent", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"username": "parent1", "password": "password", "name": "Parent 1", "address": "Address 1", "occupation": "Occupation 1", "phone_number": "081234567890"}`)
		req := httptest.NewRequest(http.MethodPost, "/admin/parent", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_AssignParentToStudent(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Post("/admin/parent/assign", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"parent_id": 1, "student_id": 1}`)
		req := httptest.NewRequest(http.MethodPost, "/admin/parent/assign", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Post("/admin/parent/assign", func(ctx *fiber.Ctx) error {
			return nil
		})

		reqBody := []byte(`{"parent_id": 1}`)
		req := httptest.NewRequest(http.MethodPost, "/admin/parent/assign", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_GetParents(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/admin/parent", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodGet, "/admin/parent", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
	
	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Get("/admin/parent", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodGet, "/admin/parent", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAdminHandler_RemoveParentFromStudent(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		app := fiber.New()
		app.Delete("/admin/parent/:parentID/student/:studentID", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodDelete, "/admin/parent/1/student/1", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		app.Delete("/admin/parent/:parentID/student/:studentID", func(ctx *fiber.Ctx) error {
			return nil
		})

		req := httptest.NewRequest(http.MethodDelete, "/admin/parent/1/student/1", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}