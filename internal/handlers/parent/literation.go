package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *ParentHandler) GetStudentLiterations(ctx *fiber.Ctx) (err error) {
	userID := ctx.Locals("user").(string)

	literations, err := h.parentService.GetStudentLiterations(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
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
			Point:          literation.Points,
			Status:         literation.Status,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    res,
	})
}

func (h *ParentHandler) GetStudentLiterationDetail(ctx *fiber.Ctx) (err error) {
	literationID := ctx.Params("literationID")

	literation, err := h.parentService.GetStudentLiterationDetail(literationID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
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
		Point:          literation.Points,
		Status:         literation.Status,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    res,
	})
}
