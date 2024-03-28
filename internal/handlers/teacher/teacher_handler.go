package handlers

import (
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	services "github.com/Magetan-Boyz/Backend/internal/services/teacher"
)

type TeacherHandler struct {
	teacherSvc        services.TeacherService
	middlewareManager middleware.Middleware
}

func NewTeacherHandler(teacherSvc services.TeacherService, middlewareManager middleware.Middleware) *TeacherHandler {
	return &TeacherHandler{
		teacherSvc:        teacherSvc,
		middlewareManager: middlewareManager,
	}
}
