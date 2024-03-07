package routes

import (
	controllers "github.com/Magetan-Boyz/Backend/internal/controllers/teacher"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	services "github.com/Magetan-Boyz/Backend/internal/services/teacher"
	"github.com/gofiber/fiber/v2"
)

func TeacherRoutes(router fiber.Router, teacherSvc services.TeacherService, mw *middleware.Middleware) {
	teacherCtrl := controllers.NewTeacherController(teacherSvc, *mw)

	teacherCtrlRoutes := router.Group("/teacher")

	teacherScheduleCtrlRoutes := teacherCtrlRoutes.Group("/schedule")
	teacherScheduleCtrlRoutes.Use(mw.Authenticate(), mw.Authorization("teacher"))
	teacherScheduleCtrlRoutes.Get("/today", teacherCtrl.GetTodaySchedule)
}