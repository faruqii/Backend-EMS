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

	studentTaskRoutes := studentRoutes.Group("/task")
	studentTaskRoutes.Get("", student.GetTask)
	studentTaskRoutes.Post("/:id/assignment", student.SubmitTaskAssignment)
	studentTaskRoutes.Get("/:id/assignment", student.GetAssignment)

	studentScheduleRoutes := studentRoutes.Group("/schedule")
	studentScheduleRoutes.Get("", student.GetSchedule)

	StudentQuizRoutes := studentRoutes.Group("/quiz")
	StudentQuizRoutes.Get("", student.GetQuiz)

}
