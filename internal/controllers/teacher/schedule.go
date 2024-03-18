package controllers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func (t *TeacherController) GetTodaySchedule(ctx *fiber.Ctx) error {
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
		dayOfWeekToInt := helper.WeekdayToInt(schedule.DayOfWeek)
		dayOfWeek := helper.ScheduleToDay(dayOfWeekToInt)

		teacherSchedules = append(teacherSchedules, dto.TeacherSchedule{
			SubjectName: schedule.Subject.Name,
			ClassName:   schedule.Class.Name,
			Day:         dayOfWeek,
			StartTime:   schedule.StartTime,
			EndTime:     schedule.EndTime,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": teacherSchedules,
	})
}
