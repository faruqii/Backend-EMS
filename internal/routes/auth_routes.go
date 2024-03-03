package routes

import (
	"github.com/Magetan-Boyz/Backend/internal/controllers/auth"
	"github.com/Magetan-Boyz/Backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router, authService services.AuthService) {
	authController := controllers.NewAuthController(authService)

	authControllerRoutes := router.Group("/auth")

	authControllerRoutes.Post("/login", authController.LogIn)
}
