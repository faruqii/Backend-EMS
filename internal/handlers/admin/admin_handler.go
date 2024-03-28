package handlers

import (
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	services "github.com/Magetan-Boyz/Backend/internal/services/admin"
)

type AdminHandler struct {
	adminService      services.AdminService
	middlewareManager middleware.Middleware
}

func NewAdminHandler(adminService services.AdminService, middlewareManager middleware.Middleware) *AdminHandler {
	return &AdminHandler{
		adminService:      adminService,
		middlewareManager: middlewareManager,
	}
}
