package handlers

import (
	"encoding/csv"
	"io"
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

func (c *AdminHandler) CreateStudentAccountFromCsv(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get the file",
		})
	}

	f, err := file.Open()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to open the file",
		})
	}
	defer f.Close()

	reader := csv.NewReader(f)
	_, err = reader.Read() // Skip the header row
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read the header row",
		})
	}

	var students []entities.Student
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to read a row from the CSV file",
			})
		}

		student := entities.Student{
			User: entities.User{
				Username: row[0],
				Password: row[1],
			},
			Name:        row[2],
			NISN:        row[3],
			Gender:      row[4],
			Address:     row[5],
			Birthplace:  row[6],
			Birthdate:   row[7],
			Province:    row[8],
			City:        row[9],
			BloodType:   row[10],
			Religion:    row[11],
			Phone:       row[12],
			ParentPhone: row[13],
			Email:       row[14],
		}
		students = append(students, student)
	}

	// Process all students using the individual student creation method
	for _, student := range students {
		err = c.adminService.CreateStudent(&student) // Adjust this line to match your method
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create student account",
			})
		}
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Students created successfully",
	})
}

func (c *AdminHandler) RemoveStudentFromClass(ctx *fiber.Ctx) (err error) {
	studentID := ctx.Params("studentID")

	err = c.adminService.RemoveStudentFromClass(studentID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Student removed from class successfully",
	})
}
