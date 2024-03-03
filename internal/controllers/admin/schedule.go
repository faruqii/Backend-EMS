package controllers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (c *AdminController) CreateSchedule(ctx *fiber.Ctx) (err error) {
	var req dto.CreateScheduleRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	schedule := entities.Schedule{
		ClassID:   req.ClassID,
		SubjectID: req.SubjectID,
		TeacherID: req.TeacherID,
		DayOfWeek: req.DayOfWeek,
		Duration:  req.Duration,
	}

	err = c.adminService.CreateSchedule(&schedule)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Get Schedule By ID
	scheduleResponse, err := c.adminService.GetScheduleByID(schedule.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.ScheduleResponse{
		ID:        scheduleResponse.ID,
		Class:     scheduleResponse.Class.Name,
		Subject:   scheduleResponse.Subject.Name,
		Teacher:   scheduleResponse.Teacher.Name,
		DayOfWeek: scheduleResponse.DayOfWeek,
		Duration:  scheduleResponse.Duration,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Schedule created successfully",
		"data":    response,
	})
}
