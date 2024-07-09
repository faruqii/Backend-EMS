package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (h *TeacherHandler) GetAllLiterations(ctx *fiber.Ctx) (err error) {
	classID := ctx.Query("classID") // Get the class_id query parameter

	var literations []entities.Literation
	if classID != "" {
		literations, err = h.teacherSvc.FilterByClassID(classID)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to filter literations by class ID",
			})
		}
	} else {
		literations, err = h.teacherSvc.GetAllLiterations()
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to get all literations",
			})
		}
	}

	var res []dto.LiterationResponse
	for _, literation := range literations {
		res = append(res, dto.LiterationResponse{
			ID:             literation.ID,
			StudentID:      literation.StudentID,
			Student:        literation.Student.Name,
			StudentClassID: literation.Student.Class.ID,
			StudentClass:   literation.Student.Class.Name,
			Title:          literation.Title,
			Description:    literation.Description,
			Documents:      literation.Documents,
			Feedback:       literation.Feedback,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    res,
	})
}

func (h *TeacherHandler) UpdateLiterationFeedback(ctx *fiber.Ctx) (err error) {
	literationID := ctx.Params("id") // Get the literation_id parameter
	var req dto.UpdateLiterationRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	_, err = h.teacherSvc.UpdateLiterationFeedback(literationID, req.Feedback)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update literation feedback",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success Update Feedback",
	})
}

func (h *TeacherHandler) GetLiterationByID(ctx *fiber.Ctx) (err error) {
	literationID := ctx.Params("id") // Get the literation_id parameter

	literation, err := h.teacherSvc.GetLiterationByID(literationID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get literation by ID",
		})
	}

	res := dto.LiterationResponse{
		ID:             literation.ID,
		StudentID:      literation.StudentID,
		Student:        literation.Student.Name,
		StudentClassID: literation.Student.Class.ID,
		StudentClass:   literation.Student.Class.Name,
		Title:          literation.Title,
		Description:    literation.Description,
		Documents:      literation.Documents,
		Feedback:       literation.Feedback,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    res,
	})
}

func (h *TeacherHandler) GetLiterationByStudentID(ctx *fiber.Ctx) (err error) {
	studentID := ctx.Params("id") // Get the student_id parameter

	literations, err := h.teacherSvc.GetLiterationByStudentID(studentID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get literation by student ID",
		})
	}

	var res []dto.LiterationResponse
	for _, literation := range literations {
		res = append(res, dto.LiterationResponse{
			ID:             literation.ID,
			StudentID:      literation.StudentID,
			Student:        literation.Student.Name,
			StudentClassID: literation.Student.Class.ID,
			StudentClass:   literation.Student.Class.Name,
			Title:          literation.Title,
			Description:    literation.Description,
			Documents:      literation.Documents,
			Feedback:       literation.Feedback,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    res,
	})
}
