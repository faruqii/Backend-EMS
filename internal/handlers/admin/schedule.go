package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func (c *AdminHandler) CreateSchedule(ctx *fiber.Ctx) (err error) {
	var req dto.CreateScheduleRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// check if teacher is assigned to subject
	isAssigned, err := c.adminService.IsTeacherAssignedToSubject(req.TeacherID, req.SubjectID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !isAssigned {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Teacher is not assigned to the subject",
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

	schedule := entities.Schedule{
		ClassID:   req.ClassID,
		SubjectID: req.SubjectID,
		TeacherID: req.TeacherID,
		DayOfWeek: req.DayOfWeek,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
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

	response := dto.ScheduleResponse{
		ID:        schedule.ID,
		Class:     schedules.Class.Name,
		Subject:   schedules.Subject.Name,
		Teacher:   schedules.Teacher.Name,
		DayOfWeek: dayOfWeek,
		StartTime: schedule.StartTime,
		EndTime:   schedule.EndTime,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Schedule created successfully",
		"data":    response,
	})
}
