package handlers

import "github.com/Magetan-Boyz/Backend/internal/services"

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}
