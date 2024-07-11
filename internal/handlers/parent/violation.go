package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *ParentHandler) GetStudentViolation(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	userID := ctx.Locals("user").(string)
	violations, err := h.parentService.GetStudentViolation(userID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// return violations response
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

func (h *ParentHandler) GetViolationByID(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	id := ctx.Params("id")
	violation, err := h.parentService.GetViolationByID(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// return violation response
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
