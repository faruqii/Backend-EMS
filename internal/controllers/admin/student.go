package controllers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (c *AdminController) CreateStudent(ctx *fiber.Ctx) (err error) {

	var req dto.StudentRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	student := entities.Student{
		User: entities.User{
			Username: req.Username,
			Password: req.Password,
		},
		Name:       req.Name,
		NISN:       req.NISN,
		Address:    req.Address,
		Birthplace: req.Birthplace,
		Birthdate:  req.Birthdate,
	}

	err = c.adminService.CreateStudent(&student)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.StudentResponse{
		ID:         student.ID,
		Name:       student.Name,
		NISN:       req.NISN,
		Address:    req.Address,
		Birthplace: req.Birthplace,
		Birthdate:  req.Birthdate,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Student created successfully",
		"data":    response,
	})
}
