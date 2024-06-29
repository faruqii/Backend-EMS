package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func (h *StudentHandler) GetSchedule(ctx *fiber.Ctx) (err error) {
	userID := ctx.Locals("user").(string)

	schedules, err := h.studentService.GetScedule(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var studentSchedule []dto.ScheduleResponse
	for _, schedule := range schedules {
		dayOfWeek := helper.WeekdayToStr(schedule.DayOfWeek)

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
