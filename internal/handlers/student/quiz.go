package handlers

import (
	"math"
	"net/http"
	"strconv"

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

	response := []dto.StudentQuizResponse{}
	for _, q := range quiz {
		response = append(response, dto.StudentQuizResponse{
			ID:          q.ID,
			ClassID:     q.Class.Name,
			SubjectID:   q.Subject.Name,
			TeacherID:   q.Teacher.Name,
			Title:       q.Title,
			TypeOfQuiz:  q.TypeOfQuiz,
			Description: q.Description,
			Deadline:    q.Deadline,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get quiz",
		"data":    response,
	})

}

func (h *StudentHandler) GetQuizQuestions(ctx *fiber.Ctx) error {
	quizID := ctx.Params("quizID")
	page, _ := strconv.Atoi(ctx.Query("page", "1"))          // Default to page 1 if not specified
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize", "10")) // Default to 10 items per page if not specified

	questions, err := h.studentService.GetQuizQuestions(quizID, page, pageSize)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	totalQuestions, err := h.studentService.CountQuizQuestions(quizID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := []dto.StudentQuestionBrief{}
	for _, q := range questions {
		response = append(response, dto.StudentQuestionBrief{
			Text:    q.Text,
			Options: q.Options,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message":         "Success get quiz questions",
		"total_questions": totalQuestions,
		"total_pages":     int(math.Ceil(float64(totalQuestions) / float64(pageSize))),
		"current_page":    page,
		"page_size":       pageSize,
		"data":            response,
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

func (h *StudentHandler) GetMyQuizGrade(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user").(string)
	quizID := ctx.Params("quizID")

	quizAssignment, err := h.studentService.GetMyQuizGrade(quizID, userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.StudentQuizAssignmentResponse{
		QuizName:    quizAssignment.Quiz.Title,
		StudentName: quizAssignment.Student.Name,
		Grade:       quizAssignment.Grade,
		Status:      quizAssignment.Status,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get quiz grade",
		"data":    response,
	})
}
