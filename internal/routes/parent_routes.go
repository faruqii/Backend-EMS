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
	parentControllerRoutes.Use(mw.Authenticate(), mw.Authorization("parent"))

	parentScheduleControllerRoutes := parentControllerRoutes.Group("/schedule")
	parentScheduleControllerRoutes.Get("", parentController.GetSchedule)

	parentTaskControllerRoutes := parentControllerRoutes.Group("/task")
	parentTaskControllerRoutes.Get("", parentController.GetTask)
}
