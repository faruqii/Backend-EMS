package routes

import (
	controllers "github.com/Magetan-Boyz/Backend/internal/controllers/admin"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	services "github.com/Magetan-Boyz/Backend/internal/services/admin"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(router fiber.Router, adminSvc services.AdminService, mw *middleware.Middleware) {
	adminCtrl := controllers.NewAdminController(adminSvc, *mw)

	adminCtrlRoutes := router.Group("/admin")

	// Subject routes with middleware chaining
	subCtrlRoutes := adminCtrlRoutes.Group("/subjects")
	subCtrlRoutes.Use(mw.Authenticate(), mw.Authorization("admin")) // Apply middleware here
	subCtrlRoutes.Post("/create", adminCtrl.CreateSubject)
	subCtrlRoutes.Get("/all", adminCtrl.GetAllSubject)
	subCtrlRoutes.Post("/:id/assign-teacher", adminCtrl.AssignTeacherToSubject)
	subCtrlRoutes.Get("/:id/teachers", adminCtrl.GetTeachersBySubjectID)

	// Teacher routes with middleware chaining
	teacherCtrlRoutes := adminCtrlRoutes.Group("/teacher")
	teacherCtrlRoutes.Use(mw.Authenticate(), mw.Authorization("admin")) // Apply middleware here
	teacherCtrlRoutes.Post("/create", adminCtrl.CreateTeacher)
	teacherCtrlRoutes.Get("/all", adminCtrl.GetAllTeacher)
	teacherCtrlRoutes.Put("/:id/update-homeroom-status", adminCtrl.UpdateTeacherHomeroomStatus)
	teacherCtrlRoutes.Get("/:id/subjects", adminCtrl.GetTeacherSubjects)

	// Class routes with middleware chaining
	classCtrlRoutes := adminCtrlRoutes.Group("/class")
	classCtrlRoutes.Use(mw.Authenticate(), mw.Authorization("admin")) // Apply middleware here
	classCtrlRoutes.Post("/create", adminCtrl.CreateClass)
	classCtrlRoutes.Post("/:id/assign-homeroom-teacher", adminCtrl.AssignHomeroomTeacher)
	classCtrlRoutes.Get("/all", adminCtrl.GetAllClass)
	classCtrlRoutes.Get("/:id/schedule", adminCtrl.GetClassSchedule)

	// Schedule routes with middleware chaining
	scheduleCtrlRoutes := adminCtrlRoutes.Group("/schedule")
	scheduleCtrlRoutes.Use(mw.Authenticate(), mw.Authorization("admin")) // Apply middleware here
	scheduleCtrlRoutes.Post("/create", adminCtrl.CreateSchedule)

	// Student routes with middleware chaining
	studentCtrlRoutes := adminCtrlRoutes.Group("/student")
	studentCtrlRoutes.Use(mw.Authenticate(), mw.Authorization("admin")) // Apply middleware here
	studentCtrlRoutes.Post("/create", adminCtrl.CreateStudent)
}
