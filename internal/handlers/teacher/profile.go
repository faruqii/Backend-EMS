package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *TeacherHandler) GetMyProfile(ctx *fiber.Ctx) (err error) {
	teacherID := ctx.Locals("user").(string)

	teacher, err := h.teacherSvc.GetMyProfile(teacherID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := dto.TeacherProfileResponse{
		ID:                teacher.ID,
		UserID:            teacher.UserID,
		Username:          teacher.User.Username,
		Name:              teacher.Name,
		Email:             teacher.Email,
		IsHomeroomTeacher: teacher.IsHomeroom,
		IsCouncelor:       teacher.IsCouncelor,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    response,
	})

}
