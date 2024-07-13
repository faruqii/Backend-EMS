package handlers

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (h *TeacherHandler) InsertGrade(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	subjectID := ctx.Params("subjectID")
	if subjectID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Subject ID is required",
		})
	}

	user := ctx.Locals("user").(string)
	teacherID, err := h.teacherSvc.GetTeacherIDByUserID(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Parse request body
	var req dto.GradeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	grade := &entities.Grade{
		StudentID:       req.StudentID,
		SubjectID:       subjectID,
		TeacherID:       teacherID,
		Semester:        req.Semester,
		AcademicYear:    req.AcademicYear,
		FormativeScores: req.FormativeScores,
		SummativeScores: req.SummativeScores,
		ProjectScores:   req.ProjectScores,
		FinalGrade:      req.FinalGrade,
	}

	// Insert grade
	_, err = h.teacherSvc.InsertGrade(grade)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Grade inserted successfully",
	})

}

func (h *TeacherHandler) GetGradeByID(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	gradeID := ctx.Params("gradeID")
	if gradeID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Grade ID is required",
		})
	}

	grade, err := h.teacherSvc.GetGradeByID(gradeID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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

func (h *TeacherHandler) GetAllGradeByStudentID(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	studentID := ctx.Params("studentID")
	if studentID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Student ID is required",
		})
	}

	grades, err := h.teacherSvc.GetAllGradeByStudentID(studentID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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

func (h *TeacherHandler) GetAllGradeBySubjectID(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	subjectID := ctx.Params("subjectID")
	if subjectID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Subject ID is required",
		})
	}

	grades, err := h.teacherSvc.GetAllGradeBySubjectID(subjectID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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

func (h *TeacherHandler) GetAllGrade(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	semester := ctx.Query("semester")
	academicYear := ctx.Query("academicYear")

	var grades []entities.Grade
	var err error

	if semester != "" && academicYear != "" {
		grades, err = h.teacherSvc.FilterBySemesterAndAcademicYear(semester, academicYear)
	} else if semester != "" {
		grades, err = h.teacherSvc.FilterBySemester(semester)
	} else if academicYear != "" {
		grades, err = h.teacherSvc.FilterByAcademicYear(academicYear)
	} else {
		grades, err = h.teacherSvc.GetAllGrade()
	}

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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

func (h *TeacherHandler) UpdateGrade(ctx *fiber.Ctx) error {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	gradeID := ctx.Params("gradeID")
	if gradeID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Grade ID is required",
		})
	}

	user := ctx.Locals("user").(string)
	teacherID, err := h.teacherSvc.GetTeacherIDByUserID(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var req dto.GradeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	grade, err := h.teacherSvc.GetGradeByID(gradeID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	grade.StudentID = req.StudentID
	grade.SubjectID = req.SubjectID
	grade.TeacherID = teacherID
	grade.Semester = req.Semester
	grade.AcademicYear = req.AcademicYear
	grade.FormativeScores = req.FormativeScores
	grade.SummativeScores = req.SummativeScores
	grade.ProjectScores = req.ProjectScores
	grade.FinalGrade = req.FinalGrade

	updatedGrade, err := h.teacherSvc.UpdateGrade(grade)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.GradeResponse{
		ID:              updatedGrade.ID,
		StudentID:       updatedGrade.StudentID,
		Student:         updatedGrade.Student.Name,
		SubjectID:       updatedGrade.SubjectID,
		Subject:         updatedGrade.Subject.Name,
		TeacherID:       updatedGrade.TeacherID,
		Teacher:         updatedGrade.Teacher.Name,
		Semester:        updatedGrade.Semester,
		AcademicYear:    updatedGrade.AcademicYear,
		FormativeScores: updatedGrade.FormativeScores,
		SummativeScores: updatedGrade.SummativeScores,
		ProjectScores:   updatedGrade.ProjectScores,
		FinalGrade:      updatedGrade.FinalGrade,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Grade updated successfully",
		"data":    response,
	})
}
