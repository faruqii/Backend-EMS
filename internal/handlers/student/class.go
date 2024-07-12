package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *StudentHandler) GetClass(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	userID := ctx.Locals("user").(string)

	class, err := h.studentService.MyClass(userID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.ClassResponse{
		ID:              class.ID,
		Name:            class.Name,
		HomeRoomTeacher: class.HomeRoomTeacher.Name,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get class",
		"data":    response,
	})
}

func (h *StudentHandler) GetSubjects(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	userID := ctx.Locals("user").(string)

	subjects, err := h.studentService.MySubjects(userID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.ClassSubjectResponse

	for _, subject := range subjects {
		response = append(response, dto.ClassSubjectResponse{
			ClassID:     subject.Class.ID,
			ClassName:   subject.Class.Name,
			SubjectID:   subject.Subject.ID,
			SubjectName: subject.Subject.Name,
			TeacherName: subject.Teacher.Name,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get subjects",
		"data":    response,
	})
}
