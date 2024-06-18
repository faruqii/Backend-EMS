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
	adminRoutes.Use(mw.Authenticate(), mw.Authorization("admin")) // Apply middleware here

	// Subject routes
	subRoutes := adminRoutes.Group("/subjects")
	subRoutes.Post("/create", admin.CreateSubject)
	subRoutes.Get("/all", admin.GetAllSubject)
	subRoutes.Post("/:id/assign-teacher", admin.AssignTeacherToSubject)
	subRoutes.Get("/:id/teachers", admin.GetTeachersBySubjectID)

	// Teacher routes
	teacherRoutes := adminRoutes.Group("/teacher")
	teacherRoutes.Post("/create", admin.CreateTeacher)
	teacherRoutes.Get("/all", admin.GetAllTeacher)
	teacherRoutes.Put("/:id/update-homeroom-status", admin.UpdateTeacherHomeroomStatus)
	teacherRoutes.Get("/:id/subjects", admin.GetTeacherSubjects)

	// Class routes
	classRoutes := adminRoutes.Group("/class")
	classRoutes.Post("/create", admin.CreateClass)
	classRoutes.Post("/:id/assign-homeroom-teacher", admin.AssignHomeroomTeacher)
	classRoutes.Get("/all", admin.GetAllClass)
	classRoutes.Get("/:id/schedule", admin.GetClassSchedule)
	classRoutes.Post("/:id/students", admin.InsertStudentToClass)
	classRoutes.Get("/:id/students", admin.GetAllStudentsBelongToClass)
	classRoutes.Post("/:id/assign-subject", admin.AssignSubjectToClass)

	// Schedule routes
	scheduleRoutes := adminRoutes.Group("/schedule")
	scheduleRoutes.Post("/create", admin.CreateSchedule)

	// Student routes
	studentRoutes := adminRoutes.Group("/student")
	studentRoutes.Post("/create", admin.CreateStudent)
	studentRoutes.Get("/all", admin.GetAllStudents)

	// Parent routes
	parentRoutes := adminRoutes.Group("/parent")
	parentRoutes.Post("/create", admin.CreateParentAccount)
	parentRoutes.Post("/assign-student", admin.AssignParentToStudent)

	// Announcement routes
	announcementRoutes := adminRoutes.Group("/announcement")
	announcementRoutes.Post("/create", admin.CreateAnnouncement)
	announcementRoutes.Get("", admin.GetAnnouncements)
	adminRoutes.Get("/:id", admin.GetAnnouncementByID)
	adminRoutes.Put("/:id", admin.UpdateAnnouncement)
	adminRoutes.Delete("/:id", admin.DeleteAnnouncement)

}
