package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *ParentHandler) GetTask(ctx *fiber.Ctx) (err error) {
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
