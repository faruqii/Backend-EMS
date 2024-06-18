package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (t *TeacherHandler) CreateQuiz(ctx *fiber.Ctx) error {

	classID := ctx.Params("classID")
	if classID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Class ID is required",
		})
	}

	subjectID := ctx.Params("subjectID")
	if subjectID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Subject ID is required",
		})
	}

	userID := ctx.Locals("user").(string)

	teacherID, err := t.teacherSvc.GetTeacherIDByUserID(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var req dto.CreateQuizRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Convert req.Questions from []dto.QuestionRequest to []entities.Question
	var questions []entities.Question
	for _, q := range req.Questions {
		question := entities.Question{
			Text:           q.Text,
			Options:        q.Options,
			CorrectAnswer:  q.CorrectAnswer,
		}
		questions = append(questions, question)
	}

	quiz := entities.Quiz{
		ClassID:     classID,
		SubjectID:   subjectID,
		TeacherID:   teacherID,
		Title:       req.Title,
		TypeOfQuiz:  req.TypeOfQuiz,
		Description: req.Description,
		Deadline:    req.Deadline,
		Questions:   questions, // Use the converted questions
	}

	err = t.teacherSvc.CreateQuiz(&quiz, questions)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Quiz created successfully",
	})
}

func (t *TeacherHandler) GetQuiz(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user").(string)

	quiz, err := t.teacherSvc.GetQuizByTeacherID(userID)
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
				Options:        question.Options,
				CorrectAnswer:  question.CorrectAnswer,
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

func (t *TeacherHandler) GetAllQuizAssignment(ctx *fiber.Ctx) error {
	quizID := ctx.Params("quizID")
	if quizID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Quiz ID is required",
		})
	}

	quizAssignment, err := t.teacherSvc.GetAllQuizAssignment(quizID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := []dto.StudentQuizAssignmentResponse{}
	for _, qa := range quizAssignment {
		response = append(response, dto.StudentQuizAssignmentResponse{
			ID:          qa.ID,
			QuizName:    qa.Quiz.Title,
			StudentName: qa.Student.Name,
			NISN:        qa.Student.NISN,
			Grade:       qa.Grade,
			Status:      qa.Status,
			SubmitAt:    qa.SubmitAt,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get all quiz assignment",
		"data":    response,
	})
}
