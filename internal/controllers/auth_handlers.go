package controllers

import "github.com/Magetan-Boyz/Backend/internal/services"

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}
