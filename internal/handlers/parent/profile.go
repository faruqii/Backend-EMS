package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *ParentHandler) GetMyProfile(ctx *fiber.Ctx) (err error) {
	userID := ctx.Locals("user").(string)

	parent, err := h.parentService.GetMyProfile(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := dto.ParentProfileResponse{
		ID:          parent.ID,
		UserID:      parent.UserID,
		Username:    parent.User.Username,
		Name:        parent.Name,
		Address:     parent.Address,
		Occupation:  parent.Occupation,
		PhoneNumber: parent.PhoneNumber,
		Email:       parent.Email,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    response,
	})

}
