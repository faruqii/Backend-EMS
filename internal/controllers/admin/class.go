package controllers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func (c *AdminController) CreateClass(ctx *fiber.Ctx) (err error) {

	var req dto.CreateClassRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	class := entities.Class{
		Name: req.Name,
	}

	err = c.adminService.CreateClass(&class)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.ClassResponse{
		ID:              class.ID,
		Name:            class.Name,
		HomeRoomTeacher: class.HomeRoomTeacher.Name,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Class created successfully",
		"data":    response,
	})
}

func (c *AdminController) AssignHomeroomTeacher(ctx *fiber.Ctx) (err error) {
	classID := ctx.Params("id")

	class, err := c.adminService.FindClassByID(classID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var req dto.AssignHomeroomTeacherRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	teacher, err := c.adminService.FindTeacherByID(req.TeacherID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = c.adminService.AssignHomeroomTeacher(classID, req.TeacherID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.ClassResponse{
		ID:              class.ID,
		Name:            class.Name,
		HomeRoomTeacher: teacher.Name,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Homeroom teacher assigned successfully",
		"data":    response,
	})
}

func (c *AdminController) GetAllClass(ctx *fiber.Ctx) (err error) {

	classes, err := c.adminService.GetAllClass()

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.ClassResponse

	for _, class := range classes {
		classRes := dto.ClassResponse{
			ID:              class.ID,
			Name:            class.Name,
			HomeRoomTeacher: class.HomeRoomTeacher.Name,
		}
		response = append(response, classRes)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}

func (c *AdminController) UpdateTeacherHomeroomStatus(ctx *fiber.Ctx) (err error) {
	teacherID := ctx.Params("id")

	req := dto.UpdateHomeroomStatusRequest{}

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = c.adminService.UpdateTeacherHomeroomStatus(teacherID, req.Status)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Homeroom status updated successfully",
	})
}

func (c *AdminController) GetClassSchedule(ctx *fiber.Ctx) (err error) {
	classID := ctx.Params("id")

	schedules, err := c.adminService.GetClassSchedule(classID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.ScheduleResponse

	for _, schedule := range schedules {
		dayOfWeekToInt := helper.WeekdayToInt(schedule.DayOfWeek)
		dayOfWeek := helper.ScheduleToDay(dayOfWeekToInt)
		loc, _ := time.LoadLocation("Asia/Jakarta")
		startTimeFormatted := schedule.StartTime.In(loc).Format("15:04")
		endTimeFormatted := schedule.EndTime.In(loc).Format("15:04")
		scheduleRes := dto.ScheduleResponse{
			ID:        schedule.ID,
			Subject:   schedule.Subject.Name,
			Teacher:   schedule.Teacher.Name,
			DayOfWeek: dayOfWeek,
			StartTime: startTimeFormatted,
			EndTime:   endTimeFormatted,
		}
		response = append(response, scheduleRes)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}
