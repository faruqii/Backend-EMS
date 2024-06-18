package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func (h *ParentHandler) GetSchedule(ctx *fiber.Ctx) (err error) {
	userID := ctx.Locals("user").(string)

	schedules, err := h.parentService.GetScheduleByStudentID(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// response
	var studentSchedule []dto.ScheduleResponse
	for _, schedule := range schedules {
		dayOfWeekToInt := helper.WeekdayToInt(schedule.DayOfWeek)
		dayOfWeek := helper.ScheduleToDay(dayOfWeekToInt)

		studentSchedule = append(studentSchedule, dto.ScheduleResponse{
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
		"message": "success get schedule",
		"data":    studentSchedule,
	})
}
