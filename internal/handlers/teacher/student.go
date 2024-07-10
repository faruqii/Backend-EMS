package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *TeacherHandler) GetAllStudentByClass(ctx *fiber.Ctx) (err error) {
	classID := ctx.Params("classID")

	students, err := h.teacherSvc.GetAllStudentByClass(classID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := []dto.StudentResponse{}
	for _, student := range students {
		response = append(response, dto.StudentResponse{
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
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get students",
		"data":    response,
	})
}
