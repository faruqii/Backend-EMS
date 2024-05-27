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

	subjectId := ctx.Params("subjectID")
	if subjectId == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Subject ID is required",
		})
	}

	id := ctx.Locals("user").(string)

	teacherID, err := t.teacherSvc.GetTeacherIDByUserID(id)
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

	quiz := entities.Quiz{
		ClassID:     classID,
		SubjectID:   subjectId,
		TeacherID:   teacherID,
		Title:       req.Title,
		TypeOfQuiz:  req.TypeOfQuiz,
		Description: req.Description,
		Deadline:    req.Deadline,
		Questions:   req.Questions,
	}

	var questions []entities.Question
	for _, q := range req.Questions {
		question := entities.Question{
			Text:          q.Text,
			Options:       q.Options,
			CorrectAnswer: q.CorrectAnswer,
		}
		questions = append(questions, question)
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
