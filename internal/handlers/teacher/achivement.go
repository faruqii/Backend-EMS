package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *TeacherHandler) GetAllAchivement(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	achivements, err := h.teacherSvc.GetAllAchivement()
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
		"message": "success get all achivements",
		"data":    res,
	})
}

func (h *TeacherHandler) GetAchivementByID(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	achivementID := ctx.Params("id")

	achivement, err := h.teacherSvc.GetAchivementByID(achivementID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res := dto.AchivementResponse{
		ID:               achivement.ID,
		StudentID:        achivement.StudentID,
		StudentName:      achivement.Student.Name,
		Title:            achivement.Title,
		TypeOfAchivement: achivement.TypeOfAchivement,
		Participation:    achivement.Participation,
		Level:            achivement.Level,
		Evidence:         achivement.Evidence,
		Status:           achivement.Status,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get achivement",
		"data":    res,
	})
}

func (h *TeacherHandler) GetAllAchivementByStudentID(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	studentID := ctx.Params("id")

	achivements, err := h.teacherSvc.GetAllAchivementByStudentID(studentID)
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
		"message": "success get achivements by student id",
		"data":    res,
	})
}

func (h *TeacherHandler) UpdateAchievement(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	achievementID := ctx.Params("id")

	var req dto.UpdateAchivementRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Fetch the achievement to be updated.
	achievement, err := h.teacherSvc.GetAchivementByID(achievementID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Update only the status field of the achievement.
	achievement.Status = req.Status

	_, err = h.teacherSvc.UpdateAchievement(achievementID, achievement)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Achievement status updated successfully",
	})
}

func (h *TeacherHandler) DeleteAchivement(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	achivementID := ctx.Params("id")

	err = h.teacherSvc.DeleteAchivement(achivementID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success delete achivement",
	})
}
