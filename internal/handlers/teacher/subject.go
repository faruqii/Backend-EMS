package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (h *TeacherHandler) CountStudent(ctx *fiber.Ctx) (err error) {
	classID := ctx.Params("classID")
	subjectID := ctx.Params("subjectID")

	students, err := h.teacherSvc.CountStudent(classID, subjectID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get students",
		"data":    students,
	})
}

func (h *TeacherHandler) GetMySubjects(ctx *fiber.Ctx) (err error) {
	userID := ctx.Locals("user").(string)

	subjects, err := h.teacherSvc.GetMySubjects(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get subjects",
		"data":    subjects,
	})
}

func (h *TeacherHandler) CreateSubjectMatter(ctx *fiber.Ctx) (err error) {
	subjectID := ctx.Params("subjectID")

	var req dto.SubjectMattterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var content []entities.SubjectMatterContent
	for _, c := range req.Content {
		content = append(content, entities.SubjectMatterContent{
			Title:       c.Title,
			Description: c.Description,
			Link:        c.Link,
		})
	}

	subjectMatter := &entities.SubjectMattter{
		SubjectID:   subjectID,
		Title:       req.Title,
		Description: req.Description,
		Content:     content,
	}

	if err := h.teacherSvc.CreateSubjectMatter(subjectMatter); err != nil {
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
		Subject:     subjectMatter.SubjectID,
		Title:       subjectMatter.Title,
		Description: subjectMatter.Description,
		Content:     responseContent,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "success create subject matter",
		"data":    response,
	})
}

func (h *TeacherHandler) GetSubjectMatterBySubjectID(ctx *fiber.Ctx) (err error) {
	subjectID := ctx.Params("subjectID")

	subjectMatters, err := h.teacherSvc.GetSubjectMatterBySubjectID(subjectID)
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
			Content:     responseContent,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get subject matters",
		"data":    response,
	})
}

func (h *TeacherHandler) GetDetailSubjectMatter(ctx *fiber.Ctx) (err error) {
	subjectMatterID := ctx.Params("subjectMatterID")

	subjectMatter, err := h.teacherSvc.GetDetailSubjectMatter(subjectMatterID)
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
		Content:     responseContent,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get subject matter",
		"data":    response,
	})
}
