package handlers

import (
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	service "github.com/Magetan-Boyz/Backend/internal/services/parent"
)

type ParentHandler struct {
	parentService     service.ParentService
	middlewareManager middleware.Middleware
}

func NewParentHandler(parentService service.ParentService, middlewareManager middleware.Middleware) *ParentHandler {
	return &ParentHandler{
		parentService:     parentService,
		middlewareManager: middlewareManager,
	}
}
