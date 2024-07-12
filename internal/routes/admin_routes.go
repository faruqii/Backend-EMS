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
	adminRoutes.Use(mw.Authenticate(), mw.Authorization("admin"), middleware.TestMode()) // Apply middleware here

	// Subject routes
	subRoutes := adminRoutes.Group("/subjects")
	subRoutes.Post("/create", admin.CreateSubject)
	subRoutes.Put("/:subjectID/update", admin.UpdateSubject)
	subRoutes.Get("/all", admin.GetAllSubject)
	subRoutes.Post("/:id/assign-teacher", admin.AssignTeacherToSubject)
	subRoutes.Get("/:id/teachers", admin.GetTeachersBySubjectID)
	// subRoutes.Get("", admin.GetTeachersByClassAndSubject)
	subRoutes.Get("/", admin.GetClassesSubjectsAndTeachers)

	// Teacher routes
	teacherRoutes := adminRoutes.Group("/teacher")
	teacherRoutes.Post("/create", admin.CreateTeacher)
	teacherRoutes.Post("/import", admin.CreateTeacherAccountFromCsv)
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
	scheduleRoutes.Get("/all", admin.GetSchedules)
	scheduleRoutes.Put("/:id/update", admin.UpdateSchedule)
	scheduleRoutes.Delete("/:id/delete", admin.DeleteSchedule)

	// Student routes
	studentRoutes := adminRoutes.Group("/student")
	studentRoutes.Post("/create", admin.CreateStudent)
	studentRoutes.Post("/import", admin.CreateStudentAccountFromCsv)
	studentRoutes.Get("/all", admin.GetAllStudents)
	studentRoutes.Post("/:id/remove-class", admin.RemoveStudentFromClass)

	// Parent routes
	parentRoutes := adminRoutes.Group("/parent")
	parentRoutes.Post("/create", admin.CreateParentAccount)
	parentRoutes.Post("/import", admin.CreateParentAccountFromCsv)
	parentRoutes.Post("/assign-student", admin.AssignParentToStudent)
	parentRoutes.Get("/all", admin.GetParents)

	// Announcement routes
	announcementRoutes := adminRoutes.Group("/announcement")
	announcementRoutes.Post("/create", admin.CreateAnnouncement)
	announcementRoutes.Get("", admin.GetAnnouncements)
	announcementRoutes.Get("/:id", admin.GetAnnouncementByID)
	announcementRoutes.Put("/:id/update", admin.UpdateAnnouncement)
	announcementRoutes.Delete("/:id/delete", admin.DeleteAnnouncement)

	// Agenda
	agendaRoutes := adminRoutes.Group("/agenda")
	agendaRoutes.Post("/create", admin.CreateAgenda)
	agendaRoutes.Get("", admin.GetAllAgenda)
	agendaRoutes.Get("/:agendaID", admin.GetAgendaByID)
	agendaRoutes.Put("/:agendaID/update", admin.UpdateAgenda)
	agendaRoutes.Delete("/:agendaID/delete", admin.DeleteAgenda)
}
