package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (c *AdminHandler) CreateStudent(ctx *fiber.Ctx) (err error) {

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

func (c *AdminHandler) GetAllStudents(ctx *fiber.Ctx) (err error) {
	students, err := c.adminService.GetAllStudents()

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.StudentResponse

	for _, student := range students {
		res := dto.StudentResponse{
			ID:         student.ID,
			Name:       student.Name,
			NISN:       student.NISN,
			Address:    student.Address,
			Birthplace: student.Birthplace,
			Birthdate:  student.Birthdate,
		}

		response = append(response, res)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}

func (c *AdminHandler) InsertStudentToClass(ctx *fiber.Ctx) (err error) {
	classID := ctx.Params("id")

	var req dto.InsertStudentToClass

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	student, err := c.adminService.InsertStudentToClass(req.StudentID, classID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.StudentClassResponse{
		ClassName:   student.Class.Name,
		StudentName: student.Name,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Student inserted to class successfully",
		"data":    response,
	})
}
