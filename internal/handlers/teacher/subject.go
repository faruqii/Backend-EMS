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

	subjectMatter := &entities.SubjectMattter{
		SubjectID:   subjectID,
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
	}

	if err := h.teacherSvc.CreateSubjectMatter(subjectMatter); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.SubjectMattterResponse{
		ID:          subjectMatter.ID,
		Subject:     subjectMatter.SubjectID,
		Title:       subjectMatter.Title,
		Description: subjectMatter.Description,
		Content:     subjectMatter.Content,
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
	for _, sm := range subjectMatters {
		res := dto.SubjectMattterResponse{
			ID:          sm.ID,
			Subject:     sm.Subject.Name,
			Title:       sm.Title,
			Description: sm.Description,
			Content:     sm.Content,
		}
		response = append(response, res)
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

	response := dto.SubjectMattterResponse{
		ID:          subjectMatter.ID,
		Subject:     subjectMatter.Subject.Name,
		Title:       subjectMatter.Title,
		Description: subjectMatter.Description,
		Content:     subjectMatter.Content,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get subject matter",
		"data":    response,
	})
}
