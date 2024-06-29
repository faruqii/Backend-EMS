package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func (t *TeacherHandler) GetTodaySchedule(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(string)

	teacher, err := t.teacherSvc.GetTeacherIDByUserID(token)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	today := time.Now().Weekday()
	schedules, err := t.teacherSvc.GetTodaySchedule(teacher, today)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var teacherSchedules []dto.TeacherSchedule
	for _, schedule := range schedules {
		dayOfWeek := helper.WeekdayToStr(schedule.DayOfWeek)

		teacherSchedules = append(teacherSchedules, dto.TeacherSchedule{
			SubjectName: schedule.Subject.Name,
			ClassName:   schedule.Class.Name,
			Day:         dayOfWeek,
			StartTime:   schedule.StartTime.Format(time.TimeOnly),
			EndTime:     schedule.EndTime.Format(time.TimeOnly),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": teacherSchedules,
	})
}

func (t *TeacherHandler) GetAllTeacherSchedule(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(string)

	teacher, err := t.teacherSvc.GetTeacherIDByUserID(token)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	schedules, err := t.teacherSvc.GetAllTeacherSchedule(teacher)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var teacherSchedules []dto.TeacherSchedule
	for _, schedule := range schedules {
		dayOfWeek := helper.WeekdayToStr(schedule.DayOfWeek)

		teacherSchedules = append(teacherSchedules, dto.TeacherSchedule{
			SubjectName: schedule.Subject.Name,
			ClassName:   schedule.Class.Name,
			Day:         dayOfWeek,
			StartTime:   schedule.StartTime.Format(time.TimeOnly),
			EndTime:     schedule.EndTime.Format(time.TimeOnly),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": teacherSchedules,
	})
}
