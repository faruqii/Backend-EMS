package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *ParentHandler) GetQuizAssignment(ctx *fiber.Ctx) error {
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
			QuizName:    q.Quiz.Title,
			StudentName: q.Student.Name,
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
