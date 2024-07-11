package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *StudentHandler) MyAttedance(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	userID := ctx.Locals("user").(string)

	attedance, err := h.studentService.MyAttedance(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// response
	var studentAttedance []dto.StudentAttedanceResponse

	for _, attedance := range attedance {
		studentAttedance = append(studentAttedance, dto.StudentAttedanceResponse{
			ID:              attedance.ID,
			StudentName:     attedance.Student.Name,
			SubjectName:     attedance.Subject.Name,
			AttedanceStatus: attedance.AttendaceStatus,
			AttedanceAt:     attedance.AttendaceAt,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get attedance",
		"data":    studentAttedance,
	})
}
