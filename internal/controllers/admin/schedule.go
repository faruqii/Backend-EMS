package controllers

import (
	"net/http"
	"time"

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

	// check if schedule already exists in the class and subject combination
	exists, err := c.adminService.IsScheduleExists(req.ClassID, req.SubjectID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if exists {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Schedule already exists",
		})
	}

	// Calculate start and end times based on input hours
	startTime := time.Date(2024, time.January, 1, req.StartTime, 0, 0, 0, time.UTC)
	endTime := time.Date(2024, time.January, 1, req.EndTime, 0, 0, 0, time.UTC)

	schedule := entities.Schedule{
		ClassID:   req.ClassID,
		SubjectID: req.SubjectID,
		TeacherID: req.TeacherID,
		DayOfWeek: req.DayOfWeek,
		StartTime: startTime,
		EndTime:   endTime,
	}

	err = c.adminService.CreateSchedule(&schedule)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	schedules, err := c.adminService.GetPreloadSchedule()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.ScheduleResponse{
		ID:        schedule.ID,
		Class:     schedules.Class.Name,
		Subject:   schedules.Subject.Name,
		Teacher:   schedules.Teacher.Name,
		DayOfWeek: schedule.DayOfWeek,
		StartTime: schedule.StartTime,
		EndTime:   schedule.EndTime,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Schedule created successfully",
		"data":    response,
	})
}
