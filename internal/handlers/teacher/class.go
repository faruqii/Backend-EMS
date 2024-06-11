package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *TeacherHandler) GetWhereIamTeachTheClass(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(string)

	class, err := h.teacherSvc.GetWhereIamTeachTheClass(user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.ClassSubjectResponse
	for _, c := range class {
		response = append(response, dto.ClassSubjectResponse{
			ClassName:   c.Class.Name,
			SubjectName: c.Subject.Name,
			TeacherName: c.Teacher.Name,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get class",
		"data":    response,
	})
}
