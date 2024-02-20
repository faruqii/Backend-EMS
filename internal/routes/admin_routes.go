package routes

import (
	"github.com/Magetan-Boyz/Backend/internal/controllers"
	"github.com/Magetan-Boyz/Backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(router fiber.Router, adminService services.AdminService) {
	adminController := controllers.NewAdminController(adminService)
	router.Post("/login", adminController.Login)
}
