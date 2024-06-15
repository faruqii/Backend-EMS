package handlers

import (
	"net/http"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (h *ParentHandler) GetGradeByID(ctx *fiber.Ctx) error {
	gradeID := ctx.Params("gradeID")

	grade, err := h.parentService.GetGradeByID(gradeID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.GradeResponse{
		ID:              grade.ID,
		StudentID:       grade.StudentID,
		Student:         grade.Student.Name,
		SubjectID:       grade.SubjectID,
		Subject:         grade.Subject.Name,
		TeacherID:       grade.TeacherID,
		Teacher:         grade.Teacher.Name,
		Semester:        grade.Semester,
		AcademicYear:    grade.AcademicYear,
		FormativeScores: grade.FormativeScores,
		SummativeScores: grade.SummativeScores,
		ProjectScores:   grade.ProjectScores,
		FinalGrade:      grade.FinalGrade,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Grade retrieved successfully",
		"data":    response,
	})
}

func (h *ParentHandler) GetStudentGrades(ctx *fiber.Ctx) (err error) {
	userID := ctx.Locals("user").(string)
	parentID, err := h.parentService.GetParentIDByUserID(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch parent",
		})
	}

	studentID, err := h.parentService.GetStudentIDByParentID(parentID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch student",
		})
	}

	semester := ctx.Query("semester")
	academicYear := ctx.Query("academic_year")

	var grades []entities.Grade

	// Filter grades based on query parameters
	if semester != "" && academicYear != "" {
		grades, err = h.parentService.FilterStudentGradesBySemesterAndAcademicYear(studentID, semester, academicYear)
	} else if semester != "" {
		grades, err = h.parentService.FilterStudentGradesBySemester(studentID, semester)
	} else if academicYear != "" {
		grades, err = h.parentService.FilterStudentGradesByAcademicYear(studentID, academicYear)
	} else {
		grades, err = h.parentService.GetAllGradeByStudentID(studentID)
	}

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := []dto.GradeResponse{}
	for _, grade := range grades {
		response = append(response, dto.GradeResponse{
			ID:              grade.ID,
			StudentID:       grade.StudentID,
			Student:         grade.Student.Name,
			SubjectID:       grade.SubjectID,
			Subject:         grade.Subject.Name,
			TeacherID:       grade.TeacherID,
			Teacher:         grade.Teacher.Name,
			Semester:        grade.Semester,
			AcademicYear:    grade.AcademicYear,
			FormativeScores: grade.FormativeScores,
			SummativeScores: grade.SummativeScores,
			ProjectScores:   grade.ProjectScores,
			FinalGrade:      grade.FinalGrade,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Grades retrieved successfully",
		"data":    response,
	})
}
