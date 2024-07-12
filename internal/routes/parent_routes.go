package routes

import (
	handlers "github.com/Magetan-Boyz/Backend/internal/handlers/parent"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	service "github.com/Magetan-Boyz/Backend/internal/services/parent"
	"github.com/gofiber/fiber/v2"
)

func ParentRoutes(router fiber.Router, teacherSvc service.ParentService, mw *middleware.Middleware) {
	parentController := handlers.NewParentHandler(teacherSvc, *mw)

	parentControllerRoutes := router.Group("/parent")
	parentControllerRoutes.Use(mw.Authenticate(), mw.Authorization("parent"), middleware.TestMode())

	parentScheduleControllerRoutes := parentControllerRoutes.Group("/schedule")
	parentScheduleControllerRoutes.Get("", parentController.GetSchedule)

	parentTaskControllerRoutes := parentControllerRoutes.Group("/task")
	parentTaskControllerRoutes.Get("", parentController.GetTask)

	parentAchievementControllerRoutes := parentControllerRoutes.Group("/achievement")
	parentAchievementControllerRoutes.Get("", parentController.GetAchivement)

	parentQuizControllerRoutes := parentControllerRoutes.Group("/quiz")
	parentQuizControllerRoutes.Get("", parentController.GetQuizAssignment)

	parentGradeControllerRoutes := parentControllerRoutes.Group("/grade")
	parentGradeControllerRoutes.Get("", parentController.GetStudentGrades)
	parentGradeControllerRoutes.Get("/:gradeID", parentController.GetGradeByID)

	parentViolationRoutes := parentControllerRoutes.Group("/violation")
	parentViolationRoutes.Get("", parentController.GetStudentViolation)
	parentViolationRoutes.Get("/:id", parentController.GetViolationByID)

	parentStudentRoutes := parentControllerRoutes.Group("/student")
	parentStudentRoutes.Get("", parentController.GetMyStudents)
	parentStudentRoutes.Get("/:studentID", parentController.GetStudentDetail)

	parentDispensationRoutes := parentControllerRoutes.Group("/dispensation")
	parentDispensationRoutes.Post("/create", parentController.CreateStudentDispensation)
	parentDispensationRoutes.Get("", parentController.GetStudentDispensations)
	parentDispensationRoutes.Get("/:dispensationID", parentController.GetStudentDispensationByID)

	parentAttendanceRoutes := parentControllerRoutes.Group("/attendance")
	parentAttendanceRoutes.Get("", parentController.GetStudentAttedance)

	parentLiterationRoutes := parentControllerRoutes.Group("/literation")
	parentLiterationRoutes.Get("", parentController.GetStudentLiterations)
	parentLiterationRoutes.Get("/:literationID", parentController.GetStudentLiterationDetail)

	parentProfileRoutes := parentControllerRoutes.Group("/profile")
	parentProfileRoutes.Get("", parentController.GetMyProfile)

}
