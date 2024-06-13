package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentGradeService interface {
	GetGradeByID(gradeID string) (*entities.Grade, error)
	GetAllGradeByStudentID(studentID string) ([]entities.Grade, error)
	FilterStudentGradesBySemester(studentID, semester string) ([]entities.Grade, error)
	FilterStudentGradesByAcademicYear(studentID, academicYear string) ([]entities.Grade, error)
	FilterStudentGradesBySemesterAndAcademicYear(studentID, semester, academicYear string) ([]entities.Grade, error)
}

func (s *studentService) GetGradeByID(gradeID string) (*entities.Grade, error) {
	grade, err := s.gradeRepo.GetGradeByID(gradeID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get grade", 500)
	}

	return grade, nil
}

func (s *studentService) GetAllGradeByStudentID(studentID string) ([]entities.Grade, error) {
	grades, err := s.gradeRepo.GetAllGradeByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get grades", 500)
	}

	return grades, nil
}

func (s *studentService) FilterStudentGradesBySemester(studentID, semester string) ([]entities.Grade, error) {
	grades, err := s.gradeRepo.FilterStudentGradesBySemester(studentID, semester)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get grades", 500)
	}

	return grades, nil
}

func (s *studentService) FilterStudentGradesByAcademicYear(studentID, academicYear string) ([]entities.Grade, error) {
	grades, err := s.gradeRepo.FilterStudentGradesByAcademicYear(studentID, academicYear)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get grades", 500)
	}

	return grades, nil
}

func (s *studentService) FilterStudentGradesBySemesterAndAcademicYear(studentID, semester, academicYear string) ([]entities.Grade, error) {
	grades, err := s.gradeRepo.FilterStudentGradesBySemesterAndAcademicYear(studentID, semester, academicYear)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get grades", 500)
	}

	return grades, nil
}
