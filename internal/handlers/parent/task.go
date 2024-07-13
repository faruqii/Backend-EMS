package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *ParentHandler) GetTask(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	userID := ctx.Locals("user").(string)

	task, err := h.parentService.GetTask(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := []dto.TaskResponse{}
	for _, t := range task {
		response = append(response, dto.TaskResponse{
			ID:          t.ID,
			ClassName:   t.Class.Name,
			SubjectName: t.Subject.Name,
			TeacherName: t.Teacher.Name,
			Title:       t.Title,
			TypeOfTask:  t.TypeOfTask,
			Description: t.Description,
			Deadline:    t.Deadline.Format(time.DateTime),
			Link:        t.Link,
		})

	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get task",
		"data":    response,
	})
}

func (h *ParentHandler) GetStudentAssignment(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	userID := ctx.Locals("user").(string)

	studentAssignment, err := h.parentService.GetStudentAssignment(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := []dto.StudentAssignmentResponse{}
	for _, sa := range studentAssignment {
		response = append(response, dto.StudentAssignmentResponse{
			ID:         sa.ID,
			Task:       sa.Task.Title,
			Student:    sa.Student.Name,
			Submission: sa.Submission,
			Grade:      sa.Grade,
			Feedback:   sa.Feedback,
			SubmitAt:   sa.SubmitAt,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get student assignment",
		"data":    response,
	})
}
