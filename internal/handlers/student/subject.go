package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *StudentHandler) GetSubjectByClassID(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	classID := ctx.Params("classID")

	subjects, err := h.studentService.GetSubjectsByClassID(classID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.ClassSubjectResponse

	for _, subject := range subjects {
		response = append(response, dto.ClassSubjectResponse{
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

func (h *StudentHandler) GetDetailSubject(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	subjectID := ctx.Params("subjectID")

	subject, err := h.studentService.GetDetailSubject(subjectID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.SubjectResponse{
		ID:          subject.ID,
		Name:        subject.Name,
		Description: subject.Description,
		Semester:    subject.Semester,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get subject",
		"data":    response,
	})

}

func (h *StudentHandler) GetSubjectMatterBySubjectID(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	subjectID := ctx.Params("subjectID")

	subjectMatters, err := h.studentService.GetSubjectMatterBySubjectID(subjectID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.SubjectMattterResponse
	for _, subjectMatter := range subjectMatters {
		responseContent := []dto.SubjectMatterContent{}
		for _, c := range subjectMatter.Content {
			responseContent = append(responseContent, dto.SubjectMatterContent{
				ID:          c.ID,
				Title:       c.Title,
				Description: c.Description,
				Link:        c.Link,
			})
		}

		response = append(response, dto.SubjectMattterResponse{
			ID:          subjectMatter.ID,
			Subject:     subjectMatter.Subject.Name,
			Title:       subjectMatter.Title,
			Description: subjectMatter.Description,
			CreatedAt:   subjectMatter.CreatedAt,
			UpdatedAt:   subjectMatter.UpdatedAt,
			Content:     responseContent,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get subject matters",
		"data":    response,
	})
}

func (h *StudentHandler) GetDetailSubjectMatter(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	subjectMatterID := ctx.Params("subjectMatterID")

	subjectMatter, err := h.studentService.GetDetailSubjectMatter(subjectMatterID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	responseContent := []dto.SubjectMatterContent{}
	for _, c := range subjectMatter.Content {
		responseContent = append(responseContent, dto.SubjectMatterContent{
			ID:          c.ID,
			Title:       c.Title,
			Description: c.Description,
			Link:        c.Link,
		})
	}

	response := dto.SubjectMattterResponse{
		ID:          subjectMatter.ID,
		Subject:     subjectMatter.Subject.Name,
		Title:       subjectMatter.Title,
		Description: subjectMatter.Description,
		CreatedAt:   subjectMatter.CreatedAt,
		UpdatedAt:   subjectMatter.UpdatedAt,
		Content:     responseContent,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get subject matter",
		"data":    response,
	})

}
