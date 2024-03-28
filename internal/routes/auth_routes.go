package routes

import (
	handlers "github.com/Magetan-Boyz/Backend/internal/handlers/auth"
	"github.com/Magetan-Boyz/Backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router, authService services.AuthService) {
	authHandler := handlers.NewAuthHandler(authService)

	authRoutes := router.Group("/auth")

	authRoutes.Post("/login", authHandler.LogIn)
}
