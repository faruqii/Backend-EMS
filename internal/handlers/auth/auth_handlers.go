package handlers

import (
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	"github.com/Magetan-Boyz/Backend/internal/services"
	service "github.com/Magetan-Boyz/Backend/internal/services/global"
)

type AuthHandler struct {
	authService       services.AuthService
	globalService     service.GlobalService
	middlewareManager middleware.Middleware
}

func NewAuthHandler(authService services.AuthService, globalService service.GlobalService, middlewareManager middleware.Middleware) *AuthHandler {
	return &AuthHandler{
		authService:       authService,
		globalService:     globalService,
		middlewareManager: middlewareManager,
	}
}
