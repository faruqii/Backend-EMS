package controllers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func (a *AuthController) LogIn(ctx *fiber.Ctx) (err error) {
	var req dto.LoginRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := a.authService.LogIn(req.Username, req.Password)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := a.authService.CreateUserToken(user, "user")
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Logged in successfully",
		"data":    response,
	})

}
