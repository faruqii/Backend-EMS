package routes

import (
	handlers "github.com/Magetan-Boyz/Backend/internal/handlers/auth"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	"github.com/Magetan-Boyz/Backend/internal/services"
	service "github.com/Magetan-Boyz/Backend/internal/services/global"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router, authService services.AuthService, globalService service.GlobalService, mw *middleware.Middleware) {
	authHandler := handlers.NewAuthHandler(authService, globalService, *mw)

	authRoutes := router.Group("/auth")

	authRoutes.Post("/login", authHandler.LogIn)

	profileRoutes := authRoutes.Group("/profile")
	profileRoutes.Use(mw.Authenticate(), mw.Authorization("user", "admin", "teacher", "student", "parent"))
	profileRoutes.Post(("change-password"), authHandler.ChangePassword)
	profileRoutes.Post(("/logout"), authHandler.LogOut)

	globalRoutes := router.Group("/global")
	globalRoutes.Get("/announcements", authHandler.GetAnnouncements)
	globalRoutes.Get("/announcements/:id", authHandler.GetAnnouncementByID)
}
