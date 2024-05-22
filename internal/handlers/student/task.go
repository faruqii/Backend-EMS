package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
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

	response := dto.TaskResponse{
		ID:          task.ID,
		ClassName:   task.Class.Name,
		SubjectName: task.Subject.Name,
		TeacherName: task.Teacher.Name,
		Title:       task.Title,
		TypeOfTask:  task.TypeOfTask,
		Description: task.Description,
		Deadline:    task.Deadline,
		Link:        task.Link,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get task",
		"data":    response,
	})
}
