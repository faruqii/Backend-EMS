package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *ParentHandler) GetAchivement(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	userID := ctx.Locals("user").(string)

	achivements, err := h.parentService.GetStudentAchievement(userID)
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
