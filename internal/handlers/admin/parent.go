package handlers

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (c *AdminHandler) CreateParentAccount(ctx *fiber.Ctx) (err error) {
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

	response := dto.ParentStudentResponse{
		ParentID:  req.ParentID,
		StudentID: req.StudentID,
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Parent assigned to student successfully",
		"data":    response,
	})
}
