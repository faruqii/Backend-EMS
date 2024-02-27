package controllers_test

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/Magetan-Boyz/Backend/internal/controllers"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	"github.com/Magetan-Boyz/Backend/internal/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"go.uber.org/mock/gomock"
)

func TestAdminController_CreateSubject(t *testing.T) {
	tests := []struct {
		name            string
		mockSvc         func(*gomock.Controller) *mocks.MockAdminService
		wantHTTPErrCode int
	}{
		{
			name: "Positive",
			mockSvc: func(ctrl *gomock.Controller) *mocks.MockAdminService {
				mockAdminService := mocks.NewMockAdminService(ctrl)
				mockAdminService.EXPECT().CreateSubject(gomock.Any()).Return(nil).Times(1)
				return mockAdminService
			},
		},
		{
			name: "Negative",
			mockSvc: func(ctrl *gomock.Controller) *mocks.MockAdminService {
				mockAdminService := mocks.NewMockAdminService(ctrl)
				mockAdminService.EXPECT().CreateSubject(gomock.Any()).Return(errors.New("internal server error")).Times(1)
				return mockAdminService
			},
			wantHTTPErrCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockSvc := tt.mockSvc(ctrl)

			mockTokenRepo := mocks.NewMockTokenRepository(ctrl)
			mockRoleRepo := mocks.NewMockRoleRepository(ctrl)

			token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZWY2ZDM3OWMtYmNiYy00MzBlLTg0MDMtY2M0NDYyMWYwMzI5IiwiZXhwIjoxNzA5MTA5NjM4LCJpYXQiOjE3MDkwMjMyMzh9.z6NCgWh0YiDt2VW9tl-yxqIOJTlXeGaZdij6XWVAwqI"

			mockTokenRepo.EXPECT().FindUserByToken(token).Return("mocked_user", nil).AnyTimes()
			mockTokenRepo.EXPECT().GetUserIDByToken(token).Return("mocked_user_id", nil).AnyTimes()

			mockRoleRepo.EXPECT().GetRoleNameFromID("mocked_user_id").Return("admin", nil).AnyTimes()

			middlewareManager := middleware.NewMiddleware(mockTokenRepo, mockRoleRepo)
			if middlewareManager == nil {
				t.Fatalf("Error creating middleware manager")
			}

			c := controllers.NewAdminController(mockSvc, *middlewareManager)

			// Create a new Fiber app for testing
			app := fiber.New()

			// Create a new request context
			ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
			ctx.Request().Header.Set("Authorization", token)

			// Set role in context locals
			ctx.Locals("user", "mocked_user")
			ctx.Locals("role", "admin")

			// Apply middleware to the context
			authMiddleware := middlewareManager.Authenticate()
			err := authMiddleware(ctx)
			if err != nil {
				t.Fatalf("Authentication failed: %v", err)
			}

			// Apply authorization middleware
			authMiddleware = middlewareManager.Authorization("admin")
			err = authMiddleware(ctx)
			if err != nil {
				t.Fatalf("Authorization failed: %v", err)
			}

			// Call the controller method
			err = c.CreateSubject(ctx)
			if err != nil {
				t.Fatalf("Error creating subject: %v", err)
			}

			// Assert HTTP status code
			actualStatusCode := ctx.Response().StatusCode()
			if tt.wantHTTPErrCode != 0 {
				assert.Equal(t, tt.wantHTTPErrCode, actualStatusCode)
				fmt.Printf("Test %s: Expected HTTP status code %d, got %d\n", tt.name, tt.wantHTTPErrCode, actualStatusCode)
			} else {
				expectedStatusCode := fiber.StatusCreated
				assert.Equal(t, expectedStatusCode, actualStatusCode)
				fmt.Printf("Test %s: Expected HTTP status code %d, got %d\n", tt.name, expectedStatusCode, actualStatusCode)
			}
		})
	}
}
