package handlers

import (
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	services "github.com/Magetan-Boyz/Backend/internal/services/student"
)

type StudentHandler struct {
	studentService    services.StudentService
	middlewareManager middleware.Middleware
}

func NewStudentHandler(studentService services.StudentService, middlewareManager middleware.Middleware) *StudentHandler {
	return &StudentHandler{
		studentService:    studentService,
		middlewareManager: middlewareManager,
	}
}
