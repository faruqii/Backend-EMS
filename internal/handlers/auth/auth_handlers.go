package handlers

import (
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type AuthHandler struct {
	authService       services.AuthService
	middlewareManager middleware.Middleware
}

func NewAuthHandler(authService services.AuthService, middlewareManager middleware.Middleware) *AuthHandler {
	return &AuthHandler{
		authService:       authService,
		middlewareManager: middlewareManager,
	}
}
