package controllers

import (
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	services "github.com/Magetan-Boyz/Backend/internal/services/teacher"
)

type TeacherController struct {
	teacherSvc        services.TeacherService
	middlewareManager middleware.Middleware
}

func NewTeacherController(teacherSvc services.TeacherService, middlewareManager middleware.Middleware) *TeacherController {
	return &TeacherController{
		teacherSvc:        teacherSvc,
		middlewareManager: middlewareManager,
	}
}
