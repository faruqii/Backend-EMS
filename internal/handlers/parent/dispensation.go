package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *ParentHandler) GetStudentDispensations(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	userID := ctx.Locals("user").(string)

	dispensations, err := h.parentService.GetStudentDispensations(userID)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.DispensationResponse
	for _, dispensation := range dispensations {
		response = append(response, dto.DispensationResponse{
			ID:        dispensation.ID,
			StudentID: dispensation.StudentID,
			Student:   dispensation.Student.Name,
			Reason:    dispensation.Reason,
			StartAt:   dispensation.StartAt.Format(time.DateTime),
			EndAt:     dispensation.EndAt.Format(time.DateTime),
			Document:  dispensation.Document,
			Status:    dispensation.Status,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get student dispensations",
		"data":    response,
	})

}

func (h *ParentHandler) GetStudentDispensationByID(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	dispensationID := ctx.Params("dispensationID")

	dispensation, err := h.parentService.GetStudentDispensationByID(dispensationID)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.DispensationResponse{
		ID:        dispensation.ID,
		StudentID: dispensation.StudentID,
		Student:   dispensation.Student.Name,
		Reason:    dispensation.Reason,
		StartAt:   dispensation.StartAt.Format(time.DateTime),
		EndAt:     dispensation.EndAt.Format(time.DateTime),
		Document:  dispensation.Document,
		Status:    dispensation.Status,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get student dispensation",
		"data":    response,
	})
}
