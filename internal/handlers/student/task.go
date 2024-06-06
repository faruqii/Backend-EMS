package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (h *StudentHandler) GetTask(ctx *fiber.Ctx) (err error) {
	userID := ctx.Locals("user").(string)

	task, err := h.studentService.GetTask(userID)
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
			Deadline:    t.Deadline,
			Link:        t.Link,
		})

	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get task",
		"data":    response,
	})
}

func (h *StudentHandler) SubmitTaskAssignment(ctx *fiber.Ctx) (err error) {
	taskID := ctx.Params("id")
	userID := ctx.Locals("user").(string)

	studentID, err := h.studentService.GetStudentIDByUserID(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var req dto.StudentAssignmentRequest
	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	asssignment := entities.StudentAssignment{
		TaskID:     taskID,
		StudentID:  studentID,
		Submission: req.Submission,
		Grade:      0,                             // default value
		Feedback:   "Menunggu untuk dinilai guru", // default value
		SubmitAt:   time.Now(),
	}

	err = h.studentService.SubmitAssignment(&asssignment)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.StudentAssignmentResponse{
		ID:         asssignment.ID,
		Task:       asssignment.TaskID,
		Student:    asssignment.StudentID,
		Submission: asssignment.Submission,
		Grade:      asssignment.Grade,
		Feedback:   asssignment.Feedback,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success submit assignment",
		"data":    response,
	})

}

func (h *StudentHandler) GetAssignment(ctx *fiber.Ctx) (err error) {
	taskID := ctx.Params("id")

	assignment, err := h.studentService.GetAssignment(taskID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.StudentAssignmentResponse{
		ID:         assignment.ID,
		Task:       assignment.Task.Title,
		Student:    assignment.Student.Name,
		Submission: assignment.Submission,
		Grade:      assignment.Grade,
		Feedback:   assignment.Feedback,
		SubmitAt:   assignment.SubmitAt,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get assignment",
		"data":    response,
	})
}
