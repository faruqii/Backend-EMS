package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (h *StudentHandler) InsertLiteration(ctx *fiber.Ctx) (err error) {
	userID := ctx.Locals("user").(string)

	var req dto.LiterationRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	studentID, err := h.studentService.GetStudentIDByUserID(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get student id",
		})
	}

	literation := &entities.Literation{
		StudentID:   studentID,
		Title:       req.Title,
		Description: req.Description,
		Documents:   req.Documents,
		Feedback:    "Menunggu Feedback",
		Points:      0,
		Status:      "Belum Dinilai",
	}

	_, err = h.studentService.InsertLiteration(literation)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to insert literation",
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Literation inserted successfully",
	})

}

func (h *StudentHandler) GetLiterationByID(ctx *fiber.Ctx) (err error) {
	literationID := ctx.Params("id")

	literation, err := h.studentService.GetLiterationByID(literationID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get literation by id",
		})
	}

	res := dto.LiterationResponse{
		ID:          literation.ID,
		StudentID:   literation.StudentID,
		Student:     literation.Student.Name,
		Title:       literation.Title,
		Description: literation.Description,
		Documents:   literation.Documents,
		Feedback:    literation.Feedback,
		Point:       literation.Points,
		Status:      literation.Status,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Literation retrieved successfully",
		"data":    res,
	})
}

func (h *StudentHandler) GetLiterationByStudentID(ctx *fiber.Ctx) (err error) {
	userID := ctx.Locals("user").(string)

	studentID, err := h.studentService.GetStudentIDByUserID(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get student id",
		})
	}

	literations, err := h.studentService.GetLiterationByStudentID(studentID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get literation by student id",
		})
	}

	var res []dto.LiterationResponse
	for _, literation := range literations {
		res = append(res, dto.LiterationResponse{
			ID:          literation.ID,
			StudentID:   literation.StudentID,
			Student:     literation.Student.Name,
			Title:       literation.Title,
			Description: literation.Description,
			Documents:   literation.Documents,
			Feedback:    literation.Feedback,
			Point:       literation.Points,
			Status:      literation.Status,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Literation retrieved successfully",
		"data":    res,
	})
}
