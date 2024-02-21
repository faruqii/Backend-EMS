package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/Magetan-Boyz/Backend/internal/controllers"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/dto"
	"github.com/Magetan-Boyz/Backend/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestAdminController_Login(t *testing.T) {
	tests := []struct {
		name            string
		mockSvc         func(*gomock.Controller) *services.MockAdminService
		args            func() *fiber.Ctx
		wantHTTPErrCode int
	}{
		{
			name: "Positive",
			mockSvc: func(ctrl *gomock.Controller) *services.MockAdminService {
				expectedAdminResponse := &entities.Admin{
					User: entities.User{
						ID:       "123",
						Username: "testuser",
					},
				}
				mockAdminService := services.NewMockAdminService(ctrl)
				mockAdminService.EXPECT().LogIn("testuser", "testpassword").
					Return(expectedAdminResponse, nil).Times(1)

				mockAdminService.EXPECT().CreateAdminToken(expectedAdminResponse).Return("mocked-token", nil).Times(1)

				return mockAdminService
			},
			args: func() *fiber.Ctx {
				reqBody := dto.AdminLoginRequest{Username: "testuser", Password: "testpassword"}
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
			mockSvc: func(ctrl *gomock.Controller) *services.MockAdminService {
				mockAdminService := services.NewMockAdminService(ctrl)
				mockAdminService.EXPECT().LogIn("testuser", "testpassword").Return(nil, errors.New("internal server error")).Times(1)

				return mockAdminService
			},
			args: func() *fiber.Ctx {
				reqBody := dto.AdminLoginRequest{Username: "testuser", Password: "testpassword"}
				reqBodyBytes, _ := json.Marshal(reqBody)

				app := fiber.New()

				ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
				ctx.Request().Header.SetContentType("application/json")
				ctx.Request().SetBody(reqBodyBytes)

				return ctx
			},
			wantHTTPErrCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			c := controllers.NewAdminController(tt.mockSvc(ctrl))
			ctx := tt.args()

			c.Login(ctx)

			if tt.wantHTTPErrCode != 0 {
				assert.Equal(t, tt.wantHTTPErrCode, ctx.Response().StatusCode())
			} else {
				assert.Equal(t, http.StatusOK, ctx.Response().StatusCode())
			}
		})
	}
}
