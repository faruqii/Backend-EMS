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
	studentRoutes.Use(mw.Authenticate(), mw.Authorization("student"))

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

	studentScheduleRoutes := studentRoutes.Group("/schedule")
	studentScheduleRoutes.Get("", student.GetSchedule)

	studentQuizRoutes := studentRoutes.Group("/quiz")
	studentQuizRoutes.Get("", student.GetQuiz)
	studentQuizRoutes.Post("/:quizID/submit", student.SubmitQuizAnswer)

	studentAttedanceRoutes := studentRoutes.Group("/attedance")
	studentAttedanceRoutes.Get("", student.MyAttedance)

	studentAchivementRoutes := studentRoutes.Group("/achivement")
	studentAchivementRoutes.Post("/create", student.CreateAchivement)
	studentAchivementRoutes.Get("", student.GetMyAchievements)

}
