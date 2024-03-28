package routes

import (
	handlers "github.com/Magetan-Boyz/Backend/internal/handlers/admin"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	services "github.com/Magetan-Boyz/Backend/internal/services/admin"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(router fiber.Router, adminSvc services.AdminService, mw *middleware.Middleware) {
	admin := handlers.NewAdminHandler(adminSvc, *mw)

	adminRoutes := router.Group("/admin")

	// Subject routes with middleware chaining
	subRoutes := adminRoutes.Group("/subjects")
	subRoutes.Use(mw.Authenticate(), mw.Authorization("admin")) // Apply middleware here
	subRoutes.Post("/create", admin.CreateSubject)
	subRoutes.Get("/all", admin.GetAllSubject)
	subRoutes.Post("/:id/assign-teacher", admin.AssignTeacherToSubject)
	subRoutes.Get("/:id/teachers", admin.GetTeachersBySubjectID)

	// Teacher routes with middleware chaining
	teacherRoutes := adminRoutes.Group("/teacher")
	teacherRoutes.Use(mw.Authenticate(), mw.Authorization("admin")) // Apply middleware here
	teacherRoutes.Post("/create", admin.CreateTeacher)
	teacherRoutes.Get("/all", admin.GetAllTeacher)
	teacherRoutes.Put("/:id/update-homeroom-status", admin.UpdateTeacherHomeroomStatus)
	teacherRoutes.Get("/:id/subjects", admin.GetTeacherSubjects)

	// Class routes with middleware chaining
	classRoutes := adminRoutes.Group("/class")
	classRoutes.Use(mw.Authenticate(), mw.Authorization("admin")) // Apply middleware here
	classRoutes.Post("/create", admin.CreateClass)
	classRoutes.Post("/:id/assign-homeroom-teacher", admin.AssignHomeroomTeacher)
	classRoutes.Get("/all", admin.GetAllClass)
	classRoutes.Get("/:id/schedule", admin.GetClassSchedule)

	// Schedule routes with middleware chaining
	scheduleRoutes := adminRoutes.Group("/schedule")
	scheduleRoutes.Use(mw.Authenticate(), mw.Authorization("admin")) // Apply middleware here
	scheduleRoutes.Post("/create", admin.CreateSchedule)

	// Student routes with middleware chaining
	studentRoutes := adminRoutes.Group("/student")
	studentRoutes.Use(mw.Authenticate(), mw.Authorization("admin")) // Apply middleware here
	studentRoutes.Post("/create", admin.CreateStudent)
	studentRoutes.Get("/all", admin.GetAllStudents)
}
