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
		AttendaceAt:     time.Now().Local(),
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

func (t *TeacherHandler) GetAttendanceBySubjectID(ctx *fiber.Ctx) error {
	subjectID := ctx.Params("subjectID")

	attendance, err := t.teacherSvc.GetAttedanceBySubjectID(subjectID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var res []dto.AttendanceResponse
	for _, attedance := range attendance {
		res = append(res, dto.AttendanceResponse{
			ID:              attedance.ID,
			StudentID:       attedance.Student.Name,
			SubjectID:       attedance.Subject.Name,
			AttendaceStatus: attedance.AttendaceStatus,
			AttendaceAt:     attedance.AttendaceAt,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get attendance",
		"data":    res,
	})
}

func (t *TeacherHandler) GetAttendanceByClassID(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(string)
	classID := ctx.Params("classID")

	// check if teacher is home room teacher for this class
	isHomeroom, err := t.teacherSvc.IsIamHomeroomTeacher(user, classID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !isHomeroom {
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": "You are not homeroom teacher for this class",
		})
	}

	attendance, err := t.teacherSvc.GetAttedanceByClassID(classID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var res []dto.AttendanceResponse
	for _, attedance := range attendance {
		res = append(res, dto.AttendanceResponse{
			ID:              attedance.ID,
			StudentID:       attedance.Student.Name,
			SubjectID:       attedance.Subject.Name,
			AttendaceStatus: attedance.AttendaceStatus,
			AttendaceAt:     attedance.AttendaceAt,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get attendance",
		"data":    res,
	})
}
