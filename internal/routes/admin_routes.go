package routes

import (
	"github.com/Magetan-Boyz/Backend/internal/controllers"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	"github.com/Magetan-Boyz/Backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(router fiber.Router, adminSvc services.AdminService, mw *middleware.Middleware) {
	adminCtrl := controllers.NewAdminController(adminSvc, *mw)

	adminCtrlRoutes := router.Group("/admin")

	// Subject routes with middleware chaining
	subCtrlRoutes := adminCtrlRoutes.Group("/subjects")
	subCtrlRoutes.Post("/create", adminCtrl.AuthAndAuthorize("admin"), adminCtrl.CreateSubject)
	subCtrlRoutes.Get("/all", adminCtrl.AuthAndAuthorize("admin"), adminCtrl.GetAllSubject)
	subCtrlRoutes.Post("/:id/assign-teacher", adminCtrl.AuthAndAuthorize("admin"), adminCtrl.AssignTeacherToSubject)
	subCtrlRoutes.Get("/:id/teachers", adminCtrl.AuthAndAuthorize("admin"), adminCtrl.GetTeachersBySubjectID)

	// Teacher routes with middleware chaining
	teacherCtrlRoutes := adminCtrlRoutes.Group("/teacher")
	teacherCtrlRoutes.Post("/create", adminCtrl.AuthAndAuthorize("admin"), adminCtrl.CreateTeacher)
	teacherCtrlRoutes.Get("/all", adminCtrl.AuthAndAuthorize("admin"), adminCtrl.GetAllTeacher)
}
