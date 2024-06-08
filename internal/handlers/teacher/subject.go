package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *TeacherHandler) CountStudent(ctx *fiber.Ctx) (err error) {
	classID := ctx.Params("classID")
	subjectID := ctx.Params("subjectID")

	students, err := h.teacherSvc.CountStudent(classID, subjectID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get students",
		"data":    students,
	})
}
