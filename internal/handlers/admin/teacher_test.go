package handlers

// import (
// 	"bytes"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"io"

// 	"github.com/Magetan-Boyz/Backend/internal/middleware"
// 	"github.com/Magetan-Boyz/Backend/internal/mocks"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/require"
// 	"go.uber.org/mock/gomock"
// )

// func TestAdminHandler_CreateTeacher(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockAdminService := mocks.NewMockAdminService(ctrl)
// 	mockMiddlewareManager := new(middleware.Middleware)

// 	handler := &AdminHandler{
// 		adminService:      mockAdminService,
// 		middlewareManager: *mockMiddlewareManager,
// 	}

// 	tests := []struct {
// 		name       string
// 		reqBody    string
// 		mockSetup  func()
// 		wantStatus int
// 		wantBody   string
// 	}{
// 		{
// 			name:    "Positive case: Successfully create teacher",
// 			reqBody: `{"username":"testuser","password":"testpass","name":"Test Name","email":"test@example.com"}`,
// 			mockSetup: func() {
// 				mockAdminService.EXPECT().CreateTeacher(gomock.Any()).Return(nil)
// 			},
// 			wantStatus: http.StatusCreated,
// 			wantBody:   `"message":"Teacher created successfully"`,
// 		},
// 		{
// 			name:    "Negative case: Body parsing error",
// 			reqBody: `invalid body`,
// 			mockSetup: func() {
// 				// No setup required as the body parsing itself will fail
// 			},
// 			wantStatus: http.StatusBadRequest,
// 			wantBody:   `"error":"invalid character 'i' looking for beginning of value"`,
// 		},
// 		{
// 			name:    "Negative case: Service error",
// 			reqBody: `{"username":"testuser","password":"testpass","name":"Test Name","email":"test@example.com"}`,
// 			mockSetup: func() {
// 				mockAdminService.EXPECT().CreateTeacher(gomock.Any()).Return(errors.New("service error"))
// 			},
// 			wantStatus: http.StatusInternalServerError,
// 			wantBody:   `"error":"service error"`,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			app := fiber.New()
// 			app.Post("/teacher", handler.CreateTeacher)

// 			tt.mockSetup()

// 			req := httptest.NewRequest(http.MethodPost, "/teacher", bytes.NewBufferString(tt.reqBody))
// 			req.Header.Set("Content-Type", "application/json")

// 			resp, err := app.Test(req, -1)
// 			require.NoError(t, err)

// 			require.Equal(t, tt.wantStatus, resp.StatusCode)
// 			body, _ := io.ReadAll(resp.Body)
// 			require.Contains(t, string(body), tt.wantBody)
// 		})
// 	}
// }
