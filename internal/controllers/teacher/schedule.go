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

		loc, _ := time.LoadLocation("Asia/Jakarta")
		startTime := schedule.StartTime.In(loc).Format("15:04:05 -0700")
		endTime := schedule.EndTime.In(loc).Format("15:04:05 -0700")

		teacherSchedules = append(teacherSchedules, dto.TeacherSchedule{
			SubjectName: schedule.Subject.Name,
			ClassName:   schedule.Class.Name,
			Day:         dayOfWeek,
			StartTime:   startTime,
			EndTime:     endTime,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": teacherSchedules,
	})
}
