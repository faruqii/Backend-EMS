package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (c *AdminHandler) CreateSubject(ctx *fiber.Ctx) (err error) {

	req := dto.SubjectRequest{}

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	subject := entities.Subject{
		Name:        req.Name,
		Description: req.Description,
		Semester:    req.Semester,
	}

	err = c.adminService.CreateSubject(&subject)

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

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Subject created successfully",
		"data":    response,
	})
}

func (c *AdminHandler) GetAllSubject(ctx *fiber.Ctx) (err error) {

	subjects, err := c.adminService.GetAllSubject()

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.SubjectResponse

	for _, subject := range subjects {
		subjectRes := dto.SubjectResponse{
			ID:          subject.ID,
			Name:        subject.Name,
			Description: subject.Description,
			Semester:    subject.Semester,
		}
		response = append(response, subjectRes)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}

func (c *AdminHandler) AssignSubjectToClass(ctx *fiber.Ctx) (err error) {
	classID := ctx.Params("id")

	var req dto.AssignSubjectToClassRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = c.adminService.AssignSubjectToClass(req.SubjectID, req.TeacherID, classID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Subject assigned to class successfully",
	})
}

func (c *AdminHandler) GetClassesSubjectsAndTeachers(ctx *fiber.Ctx) (err error) {
	classPrefix := ctx.Query("classPrefix")
	subjectID := ctx.Query("subjectID")

	if classPrefix == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "classPrefix query parameter is required",
		})
	}

	// Fetch classes based on prefix
	classes, err := c.adminService.GetClassesByPrefix(classPrefix)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Fetch subjects for these classes
	subjects, err := c.adminService.GetSubjectsByClassPrefix(classPrefix)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// If subjectID is provided, filter subjects by the subjectID
	if subjectID != "" {
		var filteredSubjects []dto.SubjectResponse
		for _, subject := range subjects {
			if subject.ID == subjectID {
				filteredSubjects = append(filteredSubjects, subject)
			}
		}
		subjects = filteredSubjects
	}

	// Fetch teachers for the specified subject if provided
	var teachers []dto.TeacherSubjectResponse
	if subjectID != "" {
		teachers, err = c.adminService.GetTeachersBySubjectID(subjectID)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	// Transform the data to match the desired format
	var data []fiber.Map
	teacherMap := make(map[string]string)

	// Create a map of subjectID to teacherName
	for _, teacher := range teachers {
		teacherMap[teacher.SubjectName] = teacher.TeacherName
	}

	for _, class := range classes {
		for _, subject := range subjects {
			teacherName := ""
			if name, ok := teacherMap[subject.Name]; ok {
				teacherName = name
			}
			entry := fiber.Map{
				"class":      class.Name,
				"class_id":   class.ID,
				"subject":    subject.Name,
				"subject_id": subject.ID,
				"teacher":    teacherName,
			}
			// Avoid adding duplicate entries
			if !contains(data, entry) {
				data = append(data, entry)
			}
		}
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Data fetched successfully",
		"data":    data,
	})
}

// Helper function to check if the entry already exists in the data slice
func contains(data []fiber.Map, entry fiber.Map) bool {
	for _, item := range data {
		if item["class_id"] == entry["class_id"] && item["subject_id"] == entry["subject_id"] {
			return true
		}
	}
	return false
}
