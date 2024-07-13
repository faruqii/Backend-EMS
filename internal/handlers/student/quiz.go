package handlers

import (
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (h *StudentHandler) GetQuiz(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

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
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	quizID := ctx.Params("quizID")
	page, _ := strconv.Atoi(ctx.Query("page", "1"))          // Default to page 1 if not specified
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize", "10")) // Default to 10 items per page if not specified

	questions, err := h.studentService.GetQuizQuestions(quizID, page, pageSize)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	typeOfQuiz, err := h.studentService.GetQuizByID(quizID)
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
		"type_of_quiz":    typeOfQuiz.TypeOfQuiz,
		"data":            response,
	})
}

func (h *StudentHandler) SubmitQuizAnswer(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

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
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	userID := ctx.Locals("user").(string)
	quizID := ctx.Params("quizID")

	quizAssignment, err := h.studentService.GetMyQuizGrade(quizID, userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.StudentQuizAssignmentResponse{
		ID:          quizAssignment.ID,
		QuizID:      quizAssignment.Quiz.ID,
		QuizName:    quizAssignment.Quiz.Title,
		StudentName: quizAssignment.Student.Name,
		NISN:        quizAssignment.Student.NISN,
		Grade:       quizAssignment.Grade,
		Status:      quizAssignment.Status,
		SubmitAt:    quizAssignment.SubmitAt.Format(time.DateTime),
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get quiz grade",
		"data":    response,
	})
}

func (h *StudentHandler) GetMyQuizGrades(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	userID := ctx.Locals("user").(string)
	subjectID := ctx.Query("subjectID")

	var assignments []entities.StudentQuizAssignment
	var err error

	if subjectID != "" {
		assignments, err = h.studentService.GetMyQuizAssignment(userID, subjectID)
	} else {
		assignments, err = h.studentService.GetMyQuizAssignment(userID, "")
	}

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := []dto.StudentQuizAssignmentResponse{}
	for _, a := range assignments {
		response = append(response, dto.StudentQuizAssignmentResponse{
			ID:          a.ID,
			QuizID:      a.Quiz.ID,
			QuizName:    a.Quiz.Title,
			StudentName: a.Student.Name,
			NISN:        a.Student.NISN,
			Grade:       a.Grade,
			Status:      a.Status,
			SubmitAt:    a.SubmitAt.Format(time.DateTime),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get quiz grades",
		"data":    response,
	})
}

func (h *StudentHandler) GetMyQuizSubmission(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	quizAssignmentID := ctx.Params("quizAssignmentID")

	quizAssignment, err := h.studentService.GetMyQuizSubmission(quizAssignmentID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.StudentQuizResponse{
		ID:          quizAssignment.ID,
		ClassID:     quizAssignment.Quiz.ClassID,
		SubjectID:   quizAssignment.Quiz.SubjectID,
		TeacherID:   quizAssignment.Quiz.TeacherID,
		Title:       quizAssignment.Quiz.Title,
		TypeOfQuiz:  quizAssignment.Quiz.TypeOfQuiz,
		Description: quizAssignment.Quiz.Description,
		Deadline:    quizAssignment.Quiz.Deadline,
	}

	questions := make([]dto.StudentQuestionBrief, len(quizAssignment.Quiz.Questions))
	for i, q := range quizAssignment.Quiz.Questions {
		questions[i] = dto.StudentQuestionBrief{
			Text:    q.Text,
			Options: q.Options,
		}
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message":   "Success get quiz submission",
		"data":      response,
		"questions": questions,
	})
}
