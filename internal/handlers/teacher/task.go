package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (t *TeacherHandler) CreateTask(ctx *fiber.Ctx) (err error) {
	token := ctx.Locals("user").(string)

	teacherID, err := t.teacherSvc.GetTeacherIDByUserID(token)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var req dto.TaskRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	task := entities.Task{
		ClassID:     req.ClassID,
		SubjectID:   req.SubjectID,
		TeacherID:   teacherID,
		Title:       req.Title,
		TypeOfTask:  req.TypeOfTask,
		Description: req.Description,
		Deadline:    req.Deadline,
		Link:        req.Link,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = t.teacherSvc.CreateTask(&task)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	taskDetails, err := t.teacherSvc.GetTask(task.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.TaskResponse{
		ID:          taskDetails.ID,
		ClassName:   taskDetails.Class.Name,
		SubjectName: taskDetails.Subject.Name,
		TeacherName: taskDetails.Teacher.Name,
		Title:       taskDetails.Title,
		TypeOfTask:  taskDetails.TypeOfTask,
		Description: taskDetails.Description,
		Deadline:    taskDetails.Deadline,
		Link:        taskDetails.Link,
		CreatedAt:   taskDetails.CreatedAt,
		UpdatedAt:   taskDetails.UpdatedAt,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Task created successfully",
		"data":    response,
	})
}
