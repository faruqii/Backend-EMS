package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (a *AuthHandler) LogIn(ctx *fiber.Ctx) (err error) {
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

	// Get user role
	role, err := a.authService.GetRoleNameFromID(user.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Check if the user is a teacher and if they are a homeroom teacher
	var isHomeroomTeacher bool
	if role == "teacher" {
		teacher, err := a.authService.GetTeacherByUserID(user.ID)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		isHomeroomTeacher = teacher.IsHomeroom
	}

	// Create token with additional claims
	token, err := a.authService.CreateUserToken(user, role, isHomeroomTeacher)
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

// ChangePassword handler
func (a *AuthHandler) ChangePassword(ctx *fiber.Ctx) (err error) {
	var req dto.ChangePasswordRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Retrieve user from context locals and deserialize it
	userID := ctx.Locals("user").(string)

	err = a.authService.ChangePassword(userID, req.OldPassword, req.NewPassword)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Password changed successfully",
	})

}

func (a *AuthHandler) LogOut(ctx *fiber.Ctx) (err error) {
	// Retrieve user from context locals and deserialize it
	userID := ctx.Locals("user").(string)

	err = a.authService.LogOut(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}
