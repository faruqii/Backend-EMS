package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func (h *TeacherHandler) GetWhereIamTeachTheClass(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	user := ctx.Locals("user").(string)

	class, err := h.teacherSvc.GetWhereIamTeachTheClass(user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.ClassSubjectResponse
	for _, c := range class {
		response = append(response, dto.ClassSubjectResponse{
			ClassID:     c.Class.ID,
			ClassName:   c.Class.Name,
			SubjectID:   c.Subject.ID,
			SubjectName: c.Subject.Name,
			TeacherName: c.Teacher.Name,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get class",
		"data":    response,
	})
}

func (h *TeacherHandler) GetStudents(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	classID := ctx.Params("classID")

	students, err := h.teacherSvc.GetAllStudents(classID)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
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
		"message": "success get students",
		"data":    response,
	})
}

func (h *TeacherHandler) GetWhereIamHomeroomTeacherinClass(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	user := ctx.Locals("user").(string)

	class, err := h.teacherSvc.GetWhereIamHomeroomTeacher(user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res := dto.ClassResponse{
		ID:              class.ID,
		Name:            class.Name,
		HomeRoomTeacher: class.HomeRoomTeacher.Name,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get class",
		"data":    res,
	})
}
