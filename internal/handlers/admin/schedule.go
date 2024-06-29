package handlers

import (
	"log"
	"net/http"
	"time"

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

	// parse in location
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// parse time
	startTime, err := time.ParseInLocation(time.TimeOnly, req.StartTime, loc)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid start time",
		})
	}

	endTime, err := time.ParseInLocation(time.TimeOnly, req.EndTime, loc)
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

	dayOfWeek := helper.WeekdayToStr(schedule.DayOfWeek)

	response := dto.ScheduleResponse{
		ID:        schedule.ID,
		Class:     schedules.Class.Name,
		Subject:   schedules.Subject.Name,
		Teacher:   schedules.Teacher.Name,
		DayOfWeek: dayOfWeek,
		StartTime: schedule.StartTime.Format(time.TimeOnly),
		EndTime:   schedule.EndTime.Format(time.TimeOnly),
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Schedule created successfully",
		"data":    response,
	})
}

func (c *AdminHandler) GetSchedules(ctx *fiber.Ctx) (err error) {
	schedules, err := c.adminService.GetAllSchedule()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.ScheduleResponse
	for _, schedule := range schedules {
		dayOfWeek := helper.WeekdayToStr(schedule.DayOfWeek)
		log.Println(dayOfWeek)

		response = append(response, dto.ScheduleResponse{
			ID:        schedule.ID,
			Class:     schedule.Class.Name,
			Subject:   schedule.Subject.Name,
			Teacher:   schedule.Teacher.Name,
			DayOfWeek: dayOfWeek,
			StartTime: schedule.StartTime.Format(time.TimeOnly),
			EndTime:   schedule.EndTime.Format(time.TimeOnly),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}

func (c *AdminHandler) UpdateSchedule(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	var req dto.CreateScheduleRequest
	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// update schedule
	schedule, err := c.adminService.GetScheduleByID(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// update schedule
	schedule.ClassID = req.ClassID
	schedule.SubjectID = req.SubjectID
	schedule.TeacherID = req.TeacherID
	schedule.DayOfWeek = req.DayOfWeek
	// parsing time
	// parse in location
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// parse time
	startTime, err := time.ParseInLocation(time.TimeOnly, req.StartTime, loc)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid start time",
		})
	}

	endTime, err := time.ParseInLocation(time.TimeOnly, req.EndTime, loc)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid end time",
		})
	}

	schedule.StartTime = startTime
	schedule.EndTime = endTime

	err = c.adminService.UpdateSchedule(schedule)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Schedule updated successfully",
	})

}

func (c *AdminHandler) DeleteSchedule(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	err = c.adminService.DeleteSchedule(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Schedule deleted successfully",
	})
}
