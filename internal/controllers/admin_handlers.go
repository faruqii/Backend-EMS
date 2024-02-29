package controllers

import (
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type AdminController struct {
	adminService      services.AdminService
	middlewareManager middleware.Middleware
}

func NewAdminController(adminService services.AdminService, middlewareManager middleware.Middleware) *AdminController {
	return &AdminController{
		adminService:      adminService,
		middlewareManager: middlewareManager,
	}
}

