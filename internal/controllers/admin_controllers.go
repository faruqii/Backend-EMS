package controllers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/dto"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func (c *AdminController) CreateSubject(ctx *fiber.Ctx) (err error) {
	authMiddleware := c.middlewareManager.Authenticate()
	err = authMiddleware(ctx)
	if err != nil {
		if middlewareErr, ok := err.(middleware.MiddlewareError); ok {
			return ctx.Status(middlewareErr.StatusCode).JSON(fiber.Map{
				"error": middlewareErr.Message,
			})
		}
	}

	// Apply authorization middleware
	authMiddleware = c.middlewareManager.Authorization("admin")
	err = authMiddleware(ctx)
	if err != nil {
		if middlewareErr, ok := err.(middleware.MiddlewareError); ok {
			return ctx.Status(middlewareErr.StatusCode).JSON(fiber.Map{
				"error": middlewareErr.Message,
			})
		}
	}

	// Proceed with the controller logic

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
	// Apply authentication middleware
	authMiddleware := c.middlewareManager.Authenticate()
	err = authMiddleware(ctx)
	if err != nil {
		return err
	}

	// Apply authorization middleware
	authMiddleware = c.middlewareManager.Authorization("admin")
	err = authMiddleware(ctx)
	if err != nil {
		return err
	}

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
		ID:    teacher.UserID,
		Name:  teacher.Name,
		Email: teacher.Email,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Teacher created successfully",
		"data":    response,
	})
}
