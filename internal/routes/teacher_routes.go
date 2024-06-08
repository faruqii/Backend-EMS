package routes

import (
	handlers "github.com/Magetan-Boyz/Backend/internal/handlers/teacher"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	services "github.com/Magetan-Boyz/Backend/internal/services/teacher"
	"github.com/gofiber/fiber/v2"
)

func TeacherRoutes(router fiber.Router, teacherService services.TeacherService, mw *middleware.Middleware) {
	teacherController := handlers.NewTeacherHandler(teacherService, *mw)

	teacherControllerRoutes := router.Group("/teacher")
	teacherControllerRoutes.Use(mw.Authenticate(), mw.Authorization("teacher"))

	teacherScheduleControllerRoutes := teacherControllerRoutes.Group("/schedule")
	teacherScheduleControllerRoutes.Get("/today", teacherController.GetTodaySchedule)
	teacherScheduleControllerRoutes.Get("/all", teacherController.GetAllTeacherSchedule)

	teacherTaskControllerRoutes := teacherControllerRoutes.Group("/task")
	teacherTaskControllerRoutes.Post("/create", teacherController.CreateTask)
	teacherTaskControllerRoutes.Get("/all", teacherController.GetAllTask)
	teacherTaskControllerRoutes.Get("/:taskID/assignment", teacherController.GetAllStudentAssignment)
	teacherTaskControllerRoutes.Put("/:assignmentID/grade", teacherController.UpdateStudentTaskAssignment)

	teacherQuizControllerRoutes := teacherControllerRoutes.Group("/quiz")
	teacherQuizControllerRoutes.Post("/:classID/:subjectID/create", teacherController.CreateQuiz)
	teacherQuizControllerRoutes.Get("", teacherController.GetQuiz)

	teacherSubjectControllerRoutes := teacherControllerRoutes.Group("/subject")
	teacherSubjectControllerRoutes.Get("/:classID/:subjectID/student", teacherController.CountStudent)

}
