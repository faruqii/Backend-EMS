package controllers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/helper"
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

	startTime, err := time.Parse("15:04", req.StartTime)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid start time",
		})
	}

	endTime, err := time.Parse("15:04", req.EndTime)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid end time",
		})
	}

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

	dayOfWeekToInt := helper.WeekdayToInt(schedule.DayOfWeek)
	dayOfWeek := helper.ScheduleToDay(dayOfWeekToInt)
	loc, _ := time.LoadLocation("Asia/Jakarta")
	startTimeFormatted := schedule.StartTime.In(loc).Format("15:04")
	endTimeFormatted := schedule.EndTime.In(loc).Format("15:04")

	response := dto.ScheduleResponse{
		ID:        schedule.ID,
		Class:     schedules.Class.Name,
		Subject:   schedules.Subject.Name,
		Teacher:   schedules.Teacher.Name,
		DayOfWeek: dayOfWeek,
		StartTime: startTimeFormatted,
		EndTime:   endTimeFormatted,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Schedule created successfully",
		"data":    response,
	})
}
