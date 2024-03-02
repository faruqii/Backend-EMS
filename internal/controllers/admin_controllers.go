package controllers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (c *AdminController) CreateSubject(ctx *fiber.Ctx) (err error) {

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

func (c *AdminController) GetAllSubject(ctx *fiber.Ctx) (err error) {

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

func (c *AdminController) CreateTeacher(ctx *fiber.Ctx) (err error) {

	var req dto.TeacherRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	teacher := entities.Teacher{
		User: entities.User{
			Username: req.Username,
			Password: req.Password,
		},
		Name:  req.Name,
		Email: req.Email,
	}

	err = c.adminService.CreateTeacher(&teacher)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.TeacherResponse{
		ID:                teacher.UserID,
		Name:              teacher.Name,
		Email:             teacher.Email,
		IsHomeroomTeacher: teacher.IsHomeroom,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Teacher created successfully",
		"data":    response,
	})
}

func (c *AdminController) GetAllTeacher(ctx *fiber.Ctx) (err error) {

	teachers, err := c.adminService.GetAllTeacher()

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.TeacherResponse

	for _, teacher := range teachers {
		teacherRes := dto.TeacherResponse{
			ID:                teacher.ID,
			Name:              teacher.Name,
			Email:             teacher.Email,
			IsHomeroomTeacher: teacher.IsHomeroom,
		}
		response = append(response, teacherRes)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}

func (c *AdminController) AssignTeacherToSubject(ctx *fiber.Ctx) (err error) {
	subjectID := ctx.Params("id")

	var req dto.TeacherSubjectRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	teacher, err := c.adminService.FindTeacherByID(req.TeacherID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	subject, err := c.adminService.FindSubjectByID(subjectID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = c.adminService.AssignTeacherToSubject(teacher.ID, subjectID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.TeacherSubjectResponse{
		SubjectName: subject.Name,
		TeacherName: teacher.Name,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Teacher assigned to subject successfully",
		"data":    response,
	})
}

func (c *AdminController) GetTeachersBySubjectID(ctx *fiber.Ctx) (err error) {
	subjectID := ctx.Params("id")

	teachers, err := c.adminService.GetTeachersBySubjectID(subjectID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Teachers fetched successfully",
		"data":    teachers,
	})
}

func (c *AdminController) CreateClass(ctx *fiber.Ctx) (err error) {

	var req dto.CreateClassRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	class := entities.Class{
		Name: req.Name,
	}

	err = c.adminService.CreateClass(&class)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.ClassResponse{
		ID:              class.ID,
		Name:            class.Name,
		HomeRoomTeacher: class.HomeRoomTeacher.Name,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Class created successfully",
		"data":    response,
	})
}

func (c *AdminController) AssignHomeroomTeacher(ctx *fiber.Ctx) (err error) {
	classID := ctx.Query("classID")
	teacherID := ctx.Query("teacherID")

	class, err := c.adminService.FindClassByID(classID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	teacher, err := c.adminService.FindTeacherByID(teacherID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = c.adminService.AssignHomeroomTeacher(classID, teacherID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.ClassResponse{
		ID:              class.ID,
		Name:            class.Name,
		HomeRoomTeacher: teacher.Name,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Homeroom teacher assigned successfully",
		"data":    response,
	})
}
