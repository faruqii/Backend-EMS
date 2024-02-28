package routes

import (
	"github.com/Magetan-Boyz/Backend/internal/controllers"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	"github.com/Magetan-Boyz/Backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(router fiber.Router, adminService services.AdminService, middlewareManager *middleware.Middleware) {
	adminController := controllers.NewAdminController(adminService, *middlewareManager)

	adminControllerRoutes := router.Group("/admin")

	// Apply middleware to subject controller routes
	subjectControllerRoutes := adminControllerRoutes.Group("/subject")
	subjectControllerRoutes.Post("/create", middlewareManager.Authenticate(), middlewareManager.Authorization("admin"), adminController.CreateSubject)

	// Apply middleware to teacher controller routes
	teacherControllerRoutes := adminControllerRoutes.Group("/teacher")
	teacherControllerRoutes.Post("/create", middlewareManager.Authenticate(), middlewareManager.Authorization("admin"), adminController.CreateTeacher)
	teacherControllerRoutes.Get("/all", middlewareManager.Authenticate(), middlewareManager.Authorization("admin"), adminController.GetAllTeacher)
}
