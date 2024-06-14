package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *TeacherHandler) GetDispenpationByID(ctx *fiber.Ctx) (err error) {
	dispensationID := ctx.Params("dispensationID")

	dispensation, err := h.teacherSvc.GetDispensationByID(dispensationID)
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
		StartAt:   dispensation.StartAt,
		EndAt:     dispensation.EndAt,
		Document:  dispensation.Document,
		Status:    dispensation.Status,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get dispensation",
		"data":    response,
	})
}

func (h *TeacherHandler) GetAllDispensations(ctx *fiber.Ctx) (err error) {
	dispensations, err := h.teacherSvc.GetAllDispensations()
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
			StartAt:   dispensation.StartAt,
			EndAt:     dispensation.EndAt,
			Document:  dispensation.Document,
			Status:    dispensation.Status,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get all dispensations",
		"data":    response,
	})
}

func (h *TeacherHandler) GetDispensationsByStudentID(ctx *fiber.Ctx) (err error) {
	studentID := ctx.Params("studentID")

	dispensations, err := h.teacherSvc.GetDispensationsByStudentID(studentID)
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
			StartAt:   dispensation.StartAt,
			EndAt:     dispensation.EndAt,
			Document:  dispensation.Document,
			Status:    dispensation.Status,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get dispensations by student id",
		"data":    response,
	})
}

func (h *TeacherHandler) UpdateDispensationStatus(ctx *fiber.Ctx) (err error) {
	dispensationID := ctx.Params("dispensationID")
	var req dto.DispensationUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, err = h.teacherSvc.UpdateDispensationStatus(dispensationID, req.Status)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success update dispensation status",
	})
}
