package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *ParentHandler) GetMyStudents(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	userID := ctx.Locals("user").(string)

	students, err := h.parentService.GetMyStudent(userID)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.StudentResponse

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
		"message": "Students fetched successfully",
		"data":    response,
	})
}

func (h *ParentHandler) GetStudentDetail(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}

	studentID := ctx.Params("studentID")

	student, err := h.parentService.GetStudentByID(studentID)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.StudentResponse{
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

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Student fetched successfully",
		"data":    response,
	})

}
