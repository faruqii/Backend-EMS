package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (h *TeacherHandler) CreateViolation(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	var req dto.ViolationRequest
	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// parse time
	startPunishment, err := time.Parse(time.DateOnly, req.StartPunishment)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid start punishment",
		})
	}

	endPunishment, err := time.Parse(time.DateOnly, req.EndPunishment)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid end punishment",
		})
	}

	violation := &entities.Violation{
		StudentID:       req.StudentID,
		SKNumber:        req.SKNumber,
		StartPunishment: startPunishment,
		EndPunishment:   endPunishment,
		Documents:       req.Documents,
		Reason:          req.Reason,
	}

	err = h.teacherSvc.CreateViolation(violation)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Violation created successfully",
	})
}

func (h *TeacherHandler) GetAllViolation(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	violations, err := h.teacherSvc.GetAllViolation()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var res []dto.ViolationResponse
	for _, v := range violations {
		res = append(res, dto.ViolationResponse{
			ID:              v.ID,
			StudentID:       v.StudentID,
			Student:         v.Student.Name,
			SKNumber:        v.SKNumber,
			StartPunishment: v.StartPunishment.Format(time.DateOnly),
			EndPunishment:   v.EndPunishment.Format(time.DateOnly),
			Documents:       v.Documents,
			Reason:          v.Reason,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    res,
	})
}

func (h *TeacherHandler) GetViolationByID(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	id := ctx.Params("id")
	violation, err := h.teacherSvc.GetViolationByID(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res := dto.ViolationResponse{
		ID:              violation.ID,
		StudentID:       violation.StudentID,
		Student:         violation.Student.Name,
		SKNumber:        violation.SKNumber,
		StartPunishment: violation.StartPunishment.Format(time.DateOnly),
		EndPunishment:   violation.EndPunishment.Format(time.DateOnly),
		Documents:       violation.Documents,
		Reason:          violation.Reason,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    res,
	})
}

func (h *TeacherHandler) GetViolationByStudentID(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	studentID := ctx.Params("student_id")
	violations, err := h.teacherSvc.GetViolationByStudentID(studentID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var res []dto.ViolationResponse
	for _, v := range violations {
		res = append(res, dto.ViolationResponse{
			ID:              v.ID,
			StudentID:       v.StudentID,
			Student:         v.Student.Name,
			SKNumber:        v.SKNumber,
			StartPunishment: v.StartPunishment.Format(time.DateOnly),
			EndPunishment:   v.EndPunishment.Format(time.DateOnly),
			Documents:       v.Documents,
			Reason:          v.Reason,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    res,
	})
}
