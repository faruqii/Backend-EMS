package handlers

import (
	"encoding/csv"
	"io"
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (c *AdminHandler) CreateParentAccount(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	var req dto.ParentRequest
	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	parent := entities.Parent{
		User: entities.User{
			Username: req.Username,
			Password: req.Password,
		},
		Name:        req.Name,
		Address:     req.Address,
		Occupation:  req.Occupation,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	}

	err = c.adminService.CreateParent(&parent)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.ParentResponse{
		ID:          parent.ID,
		Name:        parent.Name,
		Address:     parent.Address,
		Occupation:  parent.Occupation,
		PhoneNumber: parent.PhoneNumber,
		Email:       parent.Email,
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Parent created successfully",
		"data":    response,
	})

}

func (c *AdminHandler) AssignParentToStudent(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	var req dto.ParentStudentRequest
	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = c.adminService.AssignParentToStudent(req.ParentID, req.StudentID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := &dto.ParentStudentResponse{
		ParentID:  req.ParentID,
		StudentID: req.StudentID,
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Parent assigned to student successfully",
		"data":    response,
	})
}

func (c *AdminHandler) CreateParentAccountFromCsv(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

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

		parent := entities.Parent{
			User: entities.User{
				Username: row[0],
				Password: row[1],
			},
			Name:        row[2],
			Address:     row[3],
			Occupation:  row[4],
			PhoneNumber: row[5],
			Email:       row[6],
		}

		err = c.adminService.CreateParent(&parent)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create parent account",
			})
		}
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Parents created successfully",
	})
}

func (c *AdminHandler) GetParents(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	parents, err := c.adminService.GetAll()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.ParentResponse
	for _, parent := range parents {
		response = append(response, dto.ParentResponse{
			ID:          parent.ID,
			Name:        parent.Name,
			Address:     parent.Address,
			Occupation:  parent.Occupation,
			PhoneNumber: parent.PhoneNumber,
			Email:       parent.Email,
			StudentID:   parent.StudentID,
			StudentName: parent.StudentName,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}

func (c *AdminHandler) RemoveParentFromStudent(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	parentID := ctx.Params("parent_id")
	studentID := ctx.Params("student_id")

	err = c.adminService.RemoveParentFromStudent(parentID, studentID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Parent removed from student successfully",
	})
}
