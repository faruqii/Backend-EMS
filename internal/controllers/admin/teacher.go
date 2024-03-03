package controllers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

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
		ID:                teacher.ID,
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

func (c *AdminController) GetTeacherSubjects(ctx *fiber.Ctx) (err error) {
	teacherID := ctx.Params("id")

	subjects, err := c.adminService.GetTeacherSubjects(teacherID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Subjects fetched successfully",
		"data":    subjects,
	})
}
