package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (t *TeacherHandler) CreateAttendance(ctx *fiber.Ctx) error {
	subjectID := ctx.Params("subjectID")

	var req dto.CreateAttendanceRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	attendance := &entities.Atendance{
		StudentID:       req.StudentID,
		SubjectID:       subjectID,
		AttendaceStatus: req.AttendaceStatus,
		AttendaceAt:     time.Now(),
	}

	attendance, err := t.teacherSvc.CreateAttedance(attendance)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res := dto.AttendanceResponse{
		ID:              attendance.ID,
		StudentID:       attendance.StudentID,
		SubjectID:       attendance.SubjectID,
		AttendaceStatus: attendance.AttendaceStatus,
		AttendaceAt:     attendance.AttendaceAt,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Attendance created successfully",
		"data":    res,
	})
}
