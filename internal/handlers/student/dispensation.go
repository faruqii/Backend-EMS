package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (h *StudentHandler) CreateDispensation(ctx *fiber.Ctx) (err error) {
	userID := ctx.Locals("user").(string)

	var req dto.DispensationRequest
	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// parsing StartAt and EndAt to DateTime
	startAt, err := time.Parse(time.DateTime, req.StartAt)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	endAt, err := time.Parse(time.DateTime, req.EndAt)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	dispensation := &entities.Dispensation{
		Reason:   req.Reason,
		StartAt:  startAt,
		EndAt:    endAt,
		Document: req.Document,
		Status:   "pending",
	}

	_, err = h.studentService.CreateDispensation(userID, dispensation)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "success create dispensation",
	})

}

// Get Dispensation By ID
func (h *StudentHandler) GetDispensationByID(ctx *fiber.Ctx) (err error) {
	dispensationID := ctx.Params("dispensationID")

	dispensation, err := h.studentService.GetDispenpationByID(dispensationID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
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
		"message": "success get dispensation",
		"data":    response,
	})
}

// Get My Dispensations
func (h *StudentHandler) GetMyDispensations(ctx *fiber.Ctx) (err error) {
	userID := ctx.Locals("user").(string)

	dispensations, err := h.studentService.GetMyDispensations(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
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
		"message": "success get my dispensations",
		"data":    response,
	})
}
