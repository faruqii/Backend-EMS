package handlers

import (
	"net/http"
	"strings"

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
		Name:        req.Name,
		NISN:        req.NISN,
		Gender:      req.Gender,
		Address:     req.Address,
		Birthplace:  req.Birthplace,
		Birthdate:   req.Birthdate,
		Province:    req.Province,
		City:        req.City,
		BloodType:   req.BloodType,
		Religion:    req.Religion,
		Phone:       req.Phone,
		ParentPhone: req.ParentPhone,
		Email:       req.Email,
	}

	err = c.adminService.CreateStudent(&student)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Student created successfully",
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
			ID:          student.ID,
			Name:        student.Name,
			NISN:        student.NISN,
			Address:     student.Address,
			Birthplace:  student.Birthplace,
			Birthdate:   student.Birthdate,
			Gender:      student.Gender,
			Province:    student.Province,
			City:        student.City,
			BloodType:   student.BloodType,
			Religion:    student.Religion,
			Phone:       student.Phone,
			ParentPhone: student.ParentPhone,
			Email:       student.Email,
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

	_, err = c.adminService.InsertStudentToClass(req.StudentID, classID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if strings.Contains(err.Error(), "Student already in class") {
			statusCode = http.StatusBadRequest
		}
		return ctx.Status(statusCode).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Student inserted to class successfully",
	})
}
