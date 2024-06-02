package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *StudentHandler) GetQuiz(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user").(string)

	quiz, err := h.studentService.GetQuiz(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := []dto.QuizResponse{}
	for _, q := range quiz {
		var questions []dto.QuestionBrief
		for _, question := range q.Questions {
			questionBrief := dto.QuestionBrief{
				Text:    question.Text,
				Options: question.Options,
			}
			questions = append(questions, questionBrief)
		}

		response = append(response, dto.QuizResponse{
			ID:          q.ID,
			ClassID:     q.Class.Name,
			SubjectID:   q.Subject.Name,
			TeacherID:   q.Teacher.Name,
			Title:       q.Title,
			TypeOfQuiz:  q.TypeOfQuiz,
			Description: q.Description,
			Deadline:    q.Deadline,
			Questions:   questions,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get quiz",
		"data":    response,
	})

}
