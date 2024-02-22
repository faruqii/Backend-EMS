package controllers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func (c *AdminController) Login(ctx *fiber.Ctx) (err error) {
	var req dto.AdminLoginRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	admin, err := c.adminService.LogIn(req.Username, req.Password)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := c.adminService.CreateAdminToken(admin)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.AdminLoginResponse{
		ID:       admin.ID,
		Username: admin.Username,
		Token:    token,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"data":    response,
	})

}

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
			Role:     "teacher",
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
		ID:    teacher.ID,
		Name:  teacher.Name,
		Email: teacher.Email,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Teacher created successfully",
		"data":    response,
	})
}
