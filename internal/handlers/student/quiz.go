package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
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
				Text:           question.Text,
				TypeOfQuestion: question.TypeOfQuestion,
				Options:        question.Options,
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

func (h *StudentHandler) SubmitQuizAnswer(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user").(string)
	quizID := ctx.Params("quizID")

	var req dto.SubmitQuizRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	quizAssignment := &entities.StudentQuizAssignment{
		QuizID:    quizID,
		StudentID: userID,
		Answers:   req.Answers,
	}

	if err := h.studentService.SubmitQuiz(quizAssignment); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Quiz submitted successfully",
	})
}
