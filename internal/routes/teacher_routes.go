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

	teacherScheduleControllerRoutes := teacherControllerRoutes.Group("/schedule")
	teacherScheduleControllerRoutes.Use(mw.Authenticate(), mw.Authorization("teacher"))
	teacherScheduleControllerRoutes.Get("/today", teacherController.GetTodaySchedule)
}
