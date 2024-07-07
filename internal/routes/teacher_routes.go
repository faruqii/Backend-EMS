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
	teacherQuizControllerRoutes.Get("/:quizID/assignment", teacherController.GetAllQuizAssignment)
	teacherQuizControllerRoutes.Put("/:quizAssignmentID/grade", teacherController.GradeStudentQuiz)

	teacherSubjectControllerRoutes := teacherControllerRoutes.Group("/subject")
	teacherSubjectControllerRoutes.Get("/:classID/:subjectID/student", teacherController.CountStudent)
	teacherSubjectControllerRoutes.Post("/:subjectID/attendance", teacherController.CreateAttendance)
	teacherSubjectControllerRoutes.Get("/:subjectID/attendance", teacherController.GetAttendanceBySubjectID)
	teacherSubjectControllerRoutes.Put("/attendance/:attendanceID/update", teacherController.UpdateAttendance)
	teacherSubjectControllerRoutes.Get("/all", teacherController.GetMySubjects)
	teacherSubjectControllerRoutes.Post("/:subjectID/matter", teacherController.CreateSubjectMatter)
	teacherSubjectControllerRoutes.Get("/:subjectID/matter", teacherController.GetSubjectMatterBySubjectID)
	teacherSubjectControllerRoutes.Get("/matter/:subjectMatterID", teacherController.GetDetailSubjectMatter)

	teacherClassRoutes := teacherControllerRoutes.Group("/class")
	teacherClassRoutes.Get("", teacherController.GetWhereIamTeachTheClass)
	teacherClassRoutes.Get("/:classID/attendance", teacherController.GetAttendanceByClassID)
	teacherClassRoutes.Get("/:classID/students", teacherController.GetStudents)

	teacherAchivementRoutes := teacherControllerRoutes.Group("/achivement")
	teacherAchivementRoutes.Get("/all", teacherController.GetAllAchivement)
	teacherAchivementRoutes.Get("/:id", teacherController.GetAchivementByID)
	teacherAchivementRoutes.Get("/:id/all", teacherController.GetAllAchivementByStudentID)
	teacherAchivementRoutes.Put("/:id/update", teacherController.UpdateAchievement)
	teacherAchivementRoutes.Delete("/:id/delete", teacherController.DeleteAchivement)

	teacherGradeRoutes := teacherControllerRoutes.Group("/grade")
	teacherGradeRoutes.Post("/:subjectID/insert", teacherController.InsertGrade)
	teacherGradeRoutes.Get("/:gradeID", teacherController.GetGradeByID)
	teacherGradeRoutes.Get("/student/:studentID", teacherController.GetAllGradeByStudentID)
	teacherGradeRoutes.Get("/subject/:subjectID", teacherController.GetAllGradeBySubjectID)
	teacherGradeRoutes.Get("", teacherController.GetAllGrade)

	teacherDispensationRoutes := teacherControllerRoutes.Group("/dispensation")
	teacherDispensationRoutes.Get("/:dispensationID", teacherController.GetDispenpationByID)
	teacherDispensationRoutes.Get("", teacherController.GetAllDispensations)
	teacherDispensationRoutes.Get("/student/:studentID", teacherController.GetDispensationsByStudentID)
	teacherDispensationRoutes.Put("/:dispensationID/update", teacherController.UpdateDispensationStatus)

	teacherLiterationRoutes := teacherControllerRoutes.Group("/literation")
	teacherLiterationRoutes.Get("", teacherController.GetAllLiterations)
	teacherLiterationRoutes.Put("/:id/feedback", teacherController.UpdateLiterationFeedback)
	teacherLiterationRoutes.Get("/:id", teacherController.GetLiterationByID)
	teacherLiterationRoutes.Get("/student/:id", teacherController.GetLiterationByStudentID)

	teacherViolationRoutes := teacherControllerRoutes.Group("/violation")
	teacherViolationRoutes.Post("/create", teacherController.CreateViolation)
	teacherViolationRoutes.Get("/all", teacherController.GetAllViolation)
	teacherViolationRoutes.Get("/:id", teacherController.GetViolationByID)
	teacherViolationRoutes.Get("/student/:student_id", teacherController.GetViolationByStudentID)

	teacherStudentRoutes := teacherControllerRoutes.Group("/students")
	teacherStudentRoutes.Get("/:classID", teacherController.GetAllStudentByClass)

}
