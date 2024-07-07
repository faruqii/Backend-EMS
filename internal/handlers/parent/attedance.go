package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *ParentHandler) GetStudentAttedance(ctx *fiber.Ctx) (err error) {
	userID := ctx.Locals("user").(string)

	attedance, err := h.parentService.GetStudentAttedance(userID)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var studentAttedance []dto.StudentAttedanceResponse

	for _, attedance := range attedance {
		studentAttedance = append(studentAttedance, dto.StudentAttedanceResponse{
			ID:              attedance.ID,
			StudentID:       attedance.Student.ID,
			StudentName:     attedance.Student.Name,
			SubjectName:     attedance.Subject.Name,
			AttedanceStatus: attedance.AttendaceStatus,
			AttedanceAt:     attedance.AttendaceAt,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get student attedance",
		"data":    studentAttedance,
	})
}
