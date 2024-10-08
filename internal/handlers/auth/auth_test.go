package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	"github.com/Magetan-Boyz/Backend/internal/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"go.uber.org/mock/gomock"
)

func TestAuthController_LogIn(t *testing.T) {
	tests := []struct {
		name            string
		mockSvc         func(*gomock.Controller) *mocks.MockAuthService
		args            func() *fiber.Ctx
		wantHTTPErrCode int
	}{
		{
			name: "Positive",
			mockSvc: func(ctrl *gomock.Controller) *mocks.MockAuthService {
				expectedUserResponse := &entities.User{
					ID:       "123",
					Username: "testusername",
				}
				mockAuthService := mocks.NewMockAuthService(ctrl)
				mockAuthService.EXPECT().LogIn("testusername", "testpassword").Return(expectedUserResponse, nil).Times(1)
				mockAuthService.EXPECT().GetRoleNameFromID("123").Return("user", nil).Times(1)
				// apply role name from id to in create user token
				mockAuthService.EXPECT().CreateUserToken(expectedUserResponse, "user", false).Return("token", nil).Times(1)
				return mockAuthService
			},
			args: func() *fiber.Ctx {
				reqBody := dto.LoginRequest{Username: "testusername", Password: "testpassword"}
				reqBodyBytes, _ := json.Marshal(reqBody)

				app := fiber.New()

				ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
				ctx.Request().Header.SetContentType("application/json")
				ctx.Request().SetBody(reqBodyBytes)

				return ctx
			},
		},
		{
			name: "Negative",
			mockSvc: func(ctrl *gomock.Controller) *mocks.MockAuthService {
				mockAuthService := mocks.NewMockAuthService(ctrl)
				mockAuthService.EXPECT().LogIn("testusername", "testpassword").Return(nil, errors.New("Internal Server Error")).Times(1)

				return mockAuthService
			},
			args: func() *fiber.Ctx {
				reqBody := dto.LoginRequest{Username: "testusername", Password: "testpassword"}
				reqBodyBytes, _ := json.Marshal(reqBody)

				app := fiber.New()

				ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
				ctx.Request().Header.SetContentType("application/json")
				ctx.Request().SetBody(reqBodyBytes)

				return ctx
			},
			wantHTTPErrCode: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			c := NewAuthHandler(tt.mockSvc(ctrl), middleware.Middleware{})
			ctx := tt.args()

			c.LogIn(ctx)

			actualStatusCode := ctx.Response().StatusCode()

			if tt.wantHTTPErrCode != 0 {
				assert.Equal(t, tt.wantHTTPErrCode, actualStatusCode)
				fmt.Printf("Test %s: Expected HTTP status code %d, got %d\n", tt.name, tt.wantHTTPErrCode, actualStatusCode)
			} else {
				expectedStatusCode := http.StatusOK
				assert.Equal(t, expectedStatusCode, actualStatusCode)
				fmt.Printf("Test %s: Expected HTTP status code %d, got %d\n", tt.name, expectedStatusCode, actualStatusCode)
			}
		})
	}
}

func TestAuthHandler_ChangePassword(t *testing.T) {
	tests := []struct {
		name            string
		mockSvc         func(*gomock.Controller) *mocks.MockAuthService
		args            func() *fiber.Ctx
		wantHTTPErrCode int
	}{
		{
			name: "Positive",
			mockSvc: func(ctrl *gomock.Controller) *mocks.MockAuthService {
				mockAuthService := mocks.NewMockAuthService(ctrl)
				mockAuthService.EXPECT().ChangePassword("123", "oldpass", "newpassword").Return(nil).Times(1)
				return mockAuthService
			},
			args: func() *fiber.Ctx {
				reqBody := dto.ChangePasswordRequest{OldPassword: "oldpass", NewPassword: "newpassword"}
				reqBodyBytes, _ := json.Marshal(reqBody)

				app := fiber.New()

				ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
				ctx.Request().Header.SetContentType("application/json")
				ctx.Request().SetBody(reqBodyBytes)
				// Set the user ID in the context locals
				ctx.Locals("user", "123")

				return ctx
			},
			wantHTTPErrCode: http.StatusOK,
		},
		{
			name: "Negative",
			mockSvc: func(ctrl *gomock.Controller) *mocks.MockAuthService {
				mockAuthService := mocks.NewMockAuthService(ctrl)
				mockAuthService.EXPECT().ChangePassword("123", "oldpassword", "newpassword").Return(errors.New("Internal Server Error")).Times(1)

				return mockAuthService
			},
			args: func() *fiber.Ctx {
				reqBody := dto.ChangePasswordRequest{OldPassword: "oldpassword", NewPassword: "newpassword"}
				reqBodyBytes, _ := json.Marshal(reqBody)

				app := fiber.New()

				ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
				ctx.Request().Header.SetContentType("application/json")
				ctx.Request().SetBody(reqBodyBytes)
				// Set the user ID in the context locals
				ctx.Locals("user", "123")

				return ctx
			},
			wantHTTPErrCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			c := NewAuthHandler(tt.mockSvc(ctrl), middleware.Middleware{})
			ctx := tt.args()

			c.ChangePassword(ctx)

			actualStatusCode := ctx.Response().StatusCode()

			assert.Equal(t, tt.wantHTTPErrCode, actualStatusCode)
			fmt.Printf("Test %s: Expected HTTP status code %d, got %d\n", tt.name, tt.wantHTTPErrCode, actualStatusCode)
		})
	}
}

func TestAuthHandler_LogOut(t *testing.T) {
	tests := []struct {
		name            string
		mockSvc         func(*gomock.Controller) *mocks.MockAuthService
		args            func() *fiber.Ctx
		wantHTTPErrCode int
	}{
		{
			name: "Positive",
			mockSvc: func(ctrl *gomock.Controller) *mocks.MockAuthService {
				mockAuthService := mocks.NewMockAuthService(ctrl)
				mockAuthService.EXPECT().LogOut("123").Return(nil).Times(1)
				return mockAuthService
			},
			args: func() *fiber.Ctx {
				app := fiber.New()

				ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
				// Set the user ID in the context locals
				ctx.Locals("user", "123")

				return ctx
			},
			wantHTTPErrCode: http.StatusOK,
		},
		{
			name: "Negative",
			mockSvc: func(ctrl *gomock.Controller) *mocks.MockAuthService {
				mockAuthService := mocks.NewMockAuthService(ctrl)
				mockAuthService.EXPECT().LogOut("123").Return(errors.New("Internal Server Error")).Times(1)

				return mockAuthService
			},
			args: func() *fiber.Ctx {
				app := fiber.New()

				ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
				// Set the user ID in the context locals
				ctx.Locals("user", "123")

				return ctx
			},
			wantHTTPErrCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			c := NewAuthHandler(tt.mockSvc(ctrl), middleware.Middleware{})
			ctx := tt.args()

			c.LogOut(ctx)

			actualStatusCode := ctx.Response().StatusCode()

			assert.Equal(t, tt.wantHTTPErrCode, actualStatusCode)
			fmt.Printf("Test %s: Expected HTTP status code %d, got %d\n", tt.name, tt.wantHTTPErrCode, actualStatusCode)
		})
	}
}