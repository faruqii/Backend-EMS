package handlers

import (
	"fmt"
	"net/http"
	"time"

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
			Text:          q.Text,
			Options:       q.Options,
			CorrectAnswer: q.CorrectAnswer,
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
				ID:            question.ID,
				Text:          question.Text,
				Options:       question.Options,
				CorrectAnswer: question.CorrectAnswer,
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
			SubmitAt:    qa.SubmitAt.Format(time.DateOnly),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get all quiz assignment",
		"data":    response,
	})
}

func (t *TeacherHandler) GradeStudentQuiz(ctx *fiber.Ctx) error {
	quizAssignmentID := ctx.Params("quizAssignmentID")
	if quizAssignmentID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Quiz Assignment ID is required",
		})
	}

	var req dto.GradeStudentQuizRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := t.teacherSvc.GradeStudentQuiz(quizAssignmentID, req.Status, req.Grade)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success grade student quiz",
	})
}

func (t *TeacherHandler) GetStudentQuizAssignmentAnswer(ctx *fiber.Ctx) error {
	quizAssignmentID := ctx.Params("quizAssignmentID")
	if quizAssignmentID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Quiz Assignment ID is required",
		})
	}

	answers, err := t.teacherSvc.GetStudentQuizAssignmentAnswer(quizAssignmentID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := []dto.StudentQuizAssignmentAnswerResponse{}
	for _, a := range answers {
		questionsWithAnswers := []dto.QuestionWithAnswer{}
		for i, q := range a.Quiz.Questions {
			studentAnswer := ""
			if i < len(a.Answers) {
				studentAnswer = a.Answers[i]
			}
			questionsWithAnswers = append(questionsWithAnswers, dto.QuestionWithAnswer{
				Question:      q.Text,
				Answer:        studentAnswer,
				CorrectAnswer: q.CorrectAnswer,
			})
		}

		response = append(response, dto.StudentQuizAssignmentAnswerResponse{
			StudentID:   a.StudentID,
			StudentName: a.Student.Name,
			QuizID:      a.Quiz.ID,
			QuizTitle:   a.Quiz.Title,
			Questions:   questionsWithAnswers,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get student quiz assignment answer",
		"data":    response,
	})
}

func (t *TeacherHandler) UpdateQuiz(ctx *fiber.Ctx) error {
	quizID := ctx.Params("quizID")
	if quizID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Quiz ID is required",
		})
	}

	var req dto.CreateQuizRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	quiz := entities.Quiz{
		Title:       req.Title,
		TypeOfQuiz:  req.TypeOfQuiz,
		Description: req.Description,
		Deadline:    req.Deadline,
	}

	err := t.teacherSvc.UpdateQuiz(quizID, &quiz)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success update quiz",
	})
}

func (t *TeacherHandler) DeleteQuiz(ctx *fiber.Ctx) error {
	quizID := ctx.Params("quizID")
	if quizID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Quiz ID is required",
		})
	}

	err := t.teacherSvc.DeleteQuiz(quizID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success delete quiz",
	})
}

func (t *TeacherHandler) UpdateQuizQuestion(ctx *fiber.Ctx) error {
	questionID := ctx.Params("questionID")
	if questionID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Question ID is required",
		})
	}

	var req dto.QuestionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	question := entities.Question{
		Text:          req.Text,
		Options:       req.Options,
		CorrectAnswer: req.CorrectAnswer,
	}

	err := t.teacherSvc.UpdateQuestion(questionID, &question)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success update quiz question",
	})
}

func (t *TeacherHandler) AddQuestion(ctx *fiber.Ctx) error {
	quizID := ctx.Params("quizID")
	if quizID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "quizID is required",
		})
	}
	fmt.Print(quizID)

	var req dto.QuestionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// get QuizByID
	quiz, err := t.teacherSvc.GetQuiz(quizID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	question := entities.Question{
		Text:          req.Text,
		Options:       req.Options,
		CorrectAnswer: req.CorrectAnswer,
	}

	err = t.teacherSvc.AddQuestion(quiz.ID, &question)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Success add question",
	})
}

func (t *TeacherHandler) GetQuizWithQuestions(ctx *fiber.Ctx) error {
	quizID := ctx.Params("quizID")
	if quizID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Quiz ID is required",
		})
	}

	quiz, err := t.teacherSvc.GetQuizWithQuestions(quizID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	questions := []dto.QuestionBrief{}
	for _, q := range quiz.Questions {
		question := dto.QuestionBrief{
			ID:            q.ID,
			Text:          q.Text,
			Options:       q.Options,
			CorrectAnswer: q.CorrectAnswer,
		}
		questions = append(questions, question)
	}

	response := dto.QuizResponse{
		ID:          quiz.ID,
		ClassID:     quiz.Class.Name,
		SubjectID:   quiz.Subject.Name,
		TeacherID:   quiz.Teacher.Name,
		Title:       quiz.Title,
		TypeOfQuiz:  quiz.TypeOfQuiz,
		Description: quiz.Description,
		Deadline:    quiz.Deadline,
		Questions:   questions,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get quiz with questions",
		"data":    response,
	})
}

func (t *TeacherHandler) ExportQuiz(ctx *fiber.Ctx) error {
	quizID := ctx.Params("quizID")

	// Fetch quiz data
	quiz, err := t.teacherSvc.GetQuizForExport(quizID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Return as JSON file
	ctx.Response().Header.Set("Content-Disposition", "attachment; filename=quiz.json")
	return ctx.JSON(quiz)
}
