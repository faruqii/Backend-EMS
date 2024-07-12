package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
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

func (h *ParentHandler) CreateStudentDispensation(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	userID := ctx.Locals("user").(string)

	var req dto.DispensationRequest
	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// parsing StartAt and EndAt to DateTime
	startAt, err := time.Parse(time.DateOnly, req.StartAt)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	endAt, err := time.Parse(time.DateOnly, req.EndAt)
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

	_, err = h.parentService.CreateStudentDispensation(userID, dispensation)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "success create dispensation",
	})

}
