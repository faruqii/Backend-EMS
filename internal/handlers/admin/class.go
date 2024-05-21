package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func (h *AdminHandler) CreateClass(ctx *fiber.Ctx) (err error) {

	var req dto.CreateClassRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	class := entities.Class{
		Name: req.Name,
	}

	err = h.adminService.CreateClass(&class)

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

func (h *AdminHandler) AssignHomeroomTeacher(ctx *fiber.Ctx) (err error) {
	classID := ctx.Params("id")

	class, err := h.adminService.FindClassByID(classID)
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

	teacher, err := h.adminService.FindTeacherByID(req.TeacherID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = h.adminService.AssignHomeroomTeacher(classID, req.TeacherID)
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

func (h *AdminHandler) GetAllClass(ctx *fiber.Ctx) (err error) {

	classes, err := h.adminService.GetAllClass()

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

func (h *AdminHandler) UpdateTeacherHomeroomStatus(ctx *fiber.Ctx) (err error) {
	teacherID := ctx.Params("id")

	req := dto.UpdateTeacherStatusRequest{}

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = h.adminService.UpdateTeacherHomeroomStatus(teacherID, req.Status)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Homeroom status updated successfully",
	})
}

func (h *AdminHandler) GetClassSchedule(ctx *fiber.Ctx) (err error) {
	classID := ctx.Params("id")

	classExist, err := h.adminService.ClassExists(classID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !classExist {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Class not found",
		})
	}

	schedules, err := h.adminService.GetClassSchedule(classID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.ScheduleResponse

	for _, schedule := range schedules {
		dayOfWeekToInt := helper.WeekdayToInt(schedule.DayOfWeek)
		dayOfWeek := helper.ScheduleToDay(dayOfWeekToInt)
		scheduleRes := dto.ScheduleResponse{
			ID:        schedule.ID,
			Class:     schedule.Class.Name,
			Subject:   schedule.Subject.Name,
			Teacher:   schedule.Teacher.Name,
			DayOfWeek: dayOfWeek,
			StartTime: schedule.StartTime,
			EndTime:   schedule.EndTime,
		}
		response = append(response, scheduleRes)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}

func (h *AdminHandler) GetAllStudentsBelongToClass(ctx *fiber.Ctx) (err error) {
	classID := ctx.Params("id")

	students, err := h.adminService.GetAllStudentsBelongToClass(classID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.StudentClassResponse

	for _, student := range students {
		studentRes := dto.StudentClassResponse{
			Name: student.Name,
		}
		response = append(response, studentRes)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}
