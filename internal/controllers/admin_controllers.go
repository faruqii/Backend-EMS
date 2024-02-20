package controllers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func (c *AdminController) Login(ctx *fiber.Ctx) (err error) {
	var req dto.AdminLoginRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	admin, err := c.adminService.LogIn(req.Username, req.Password)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := c.adminService.CreateAdminToken(admin)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.AdminLoginResponse{
		ID:       admin.ID,
		Username: admin.Username,
		Token:    token,
	}

	return ctx.Status(http.StatusOK).JSON(response)
	
}
