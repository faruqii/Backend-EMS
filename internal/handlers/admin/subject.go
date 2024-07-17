package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (c *AdminHandler) CreateSubject(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

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
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

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
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

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
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	classPrefix := ctx.Query("classPrefix")
	subjectID := ctx.Query("subjectID")

	if classPrefix == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "classPrefix query parameter is required",
		})
	}

	// Fetch class-subject-teacher mapping
	classSubjects, err := c.adminService.GetClassSubjectsByPrefixAndSubject(classPrefix, subjectID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Check if no records are found
	if len(classSubjects) == 0 {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "No data found",
			"data":    []fiber.Map{},
		})
	}

	// Transform the data to match the desired format
	var data []fiber.Map
	for _, cs := range classSubjects {
		item := fiber.Map{
			"class_id":   cs.ClassID,
			"class_name": cs.Class.Name,
			"subject_id": cs.SubjectID,
			"subject":    cs.Subject.Name,
			"teacher_id": cs.TeacherID,
			"teacher":    cs.Teacher.Name,
		}
		data = append(data, item)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Data fetched successfully",
		"data":    data,
	})
}

func (c *AdminHandler) UpdateSubject(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	subjectID := ctx.Params("subjectID")

	var req dto.SubjectRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	subject, err := c.adminService.FindSubjectByID(subjectID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	subject.Name = req.Name
	subject.Description = req.Description
	subject.Semester = req.Semester

	err = c.adminService.UpdateSubject(subject.ID, subject)
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
		"message": "Subject updated successfully",
		"data":    response,
	})
}
