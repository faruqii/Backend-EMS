package routes

import (
	handlers "github.com/Magetan-Boyz/Backend/internal/handlers/student"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	services "github.com/Magetan-Boyz/Backend/internal/services/student"
	"github.com/gofiber/fiber/v2"
)

func StudentRoutes(router fiber.Router, studentSvc services.StudentService, mw *middleware.Middleware) {
	student := handlers.NewStudentHandler(studentSvc, *mw)

	studentRoutes := router.Group("/student")
	studentRoutes.Use(mw.Authenticate(), mw.Authorization("student"), middleware.TestMode())

	studentProfileRoutes := studentRoutes.Group("/profile")
	studentProfileRoutes.Get("", student.MyProfile)

	studentClassRoutes := studentRoutes.Group("/class")
	studentClassRoutes.Get("", student.GetClass)
	studentClassRoutes.Get("/subjects", student.GetSubjects)

	studentSubjectRoutes := studentRoutes.Group("/subject")
	studentSubjectRoutes.Get("/:classID", student.GetSubjectByClassID)
	studentSubjectRoutes.Get("/:subjectID/detail", student.GetDetailSubject)
	studentSubjectRoutes.Get("/:subjectID/matter", student.GetSubjectMatterBySubjectID)
	studentSubjectRoutes.Get("/matter/:subjectMatterID/detail", student.GetDetailSubjectMatter)

	studentTaskRoutes := studentRoutes.Group("/task")
	studentTaskRoutes.Get("", student.GetTask)
	studentTaskRoutes.Post("/:id/assignment", student.SubmitTaskAssignment)
	studentTaskRoutes.Get("/:id/assignment", student.GetAssignment)
	studentTaskRoutes.Put("/assignment/:assignmentID/update", student.UpdateTaskSubmission)

	studentScheduleRoutes := studentRoutes.Group("/schedule")
	studentScheduleRoutes.Get("", student.GetSchedule)

	studentQuizRoutes := studentRoutes.Group("/quiz")
	studentQuizRoutes.Get("", student.GetQuiz)
	studentQuizRoutes.Post("/:quizID/submit", student.SubmitQuizAnswer)
	studentQuizRoutes.Get("/:quizID/questions", student.GetQuizQuestions)
	studentQuizRoutes.Get("/:quizID/grade", student.GetMyQuizGrade)
	studentQuizRoutes.Get("/grades", student.GetMyQuizGrades)
	studentQuizRoutes.Get("/:quizAssignmentID/submission", student.GetMyQuizSubmission)

	studentAttedanceRoutes := studentRoutes.Group("/attedance")
	studentAttedanceRoutes.Get("", student.MyAttedance)

	studentAchivementRoutes := studentRoutes.Group("/achivement")
	studentAchivementRoutes.Post("/create", student.CreateAchivement)
	studentAchivementRoutes.Get("", student.GetMyAchievements)

	studentGradeRoutes := studentRoutes.Group("/grades")
	studentGradeRoutes.Get("", student.GetMyGrades)
	studentGradeRoutes.Get("/:gradeID", student.GetGradeByID)

	studentDispensationRoutes := studentRoutes.Group("/dispensation")
	studentDispensationRoutes.Post("/create", student.CreateDispensation)
	studentDispensationRoutes.Get("/:dispensationID", student.GetDispensationByID)
	studentDispensationRoutes.Get("", student.GetMyDispensations)

	studentLiterationRoutes := studentRoutes.Group("/literation")
	studentLiterationRoutes.Post("/create", student.InsertLiteration)
	studentLiterationRoutes.Get("/:id", student.GetLiterationByID)
	studentLiterationRoutes.Get("", student.GetLiterationByStudentID)

	studentViolationRoutes := studentRoutes.Group("/violation")
	studentViolationRoutes.Get("", student.GetMyViolation)
	studentViolationRoutes.Get("/:id", student.GetViolationByID)

}
