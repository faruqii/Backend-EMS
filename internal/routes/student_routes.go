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

	studentTaskRoutes := studentRoutes.Group("/task")
	studentTaskRoutes.Get("", student.GetTask)

}
