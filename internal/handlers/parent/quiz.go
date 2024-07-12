package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *ParentHandler) GetQuizAssignment(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	userID := ctx.Locals("user").(string)

	quiz, err := h.parentService.GetQuizAssignment(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := []dto.StudentQuizAssignmentResponse{}
	for _, q := range quiz {
		response = append(response, dto.StudentQuizAssignmentResponse{
			ID:          q.ID,
			QuizID:      q.Quiz.ID,
			QuizName:    q.Quiz.Title,
			StudentName: q.Student.Name,
			NISN:        q.Student.NISN,
			Grade:       q.Grade,
			Status:      q.Status,
			SubmitAt:    q.SubmitAt.Format(time.DateTime),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    response,
	})
}
