package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *TeacherHandler) GetAllAchivement(ctx *fiber.Ctx) (err error) {
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

func (h *TeacherHandler) UpdateAchivement(ctx *fiber.Ctx) (err error) {
	achivementID := ctx.Params("id")

	var req dto.UpdateAchivementRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	achivement, err := h.teacherSvc.GetAchivementByID(achivementID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	achivement.Status = req.Status

	_, err = h.teacherSvc.UpdateAchivement(achivementID, achivement)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success update achivement",
	})

}

func (h *TeacherHandler) DeleteAchivement(ctx *fiber.Ctx) (err error) {
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
