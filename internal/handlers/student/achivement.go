package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (h *StudentHandler) CreateAchivement(ctx *fiber.Ctx) (err error) {
	userID := ctx.Locals("user").(string)

	var req dto.AchivementRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	achivement := &entities.StudentAchivement{
		Title:            req.Title,
		TypeOfAchivement: req.TypeOfAchivement,
		Participation:    req.Participation,
		Level:            req.Level,
		Evidence:         req.Evidence,
		Status:           "pending",
	}

	_, err = h.studentService.CreateAchivement(userID, achivement)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "success create achivement",
	})
}

func (h *StudentHandler) GetMyAchievements(ctx *fiber.Ctx) (err error) {
	userID := ctx.Locals("user").(string)

	achivements, err := h.studentService.GetMyAchievements(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var res []dto.AchivementResponse
	for _, achivement := range achivements {
		res = append(res, dto.AchivementResponse{
			ID:               achivement.ID,
			StudentID:        achivement.StudentID,
			StudentName:      achivement.Student.Name,
			Title:            achivement.Title,
			TypeOfAchivement: achivement.TypeOfAchivement,
			Participation:    achivement.Participation,
			Level:            achivement.Level,
			Evidence:         achivement.Evidence,
			Status:           achivement.Status,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get achivements",
		"data":    res,
	})
}
