package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherGradeService interface {
	InsertGrade(grade *entities.Grade) (*entities.Grade, error)
	GetGradeByID(gradeID string) (*entities.Grade, error)
	GetAllGradeByStudentID(studentID string) ([]entities.Grade, error)
	GetAllGradeBySubjectID(subjectID string) ([]entities.Grade, error)
	GetAllGrade() ([]entities.Grade, error)
	FilterBySemester(semester string) ([]entities.Grade, error)
	FilterByAcademicYear(academicYear string) ([]entities.Grade, error)
	FilterBySemesterAndAcademicYear(semester, academicYear string) ([]entities.Grade, error)
	UpdateGrade(grade *entities.Grade) (*entities.Grade, error)
}

func (s *teacherService) InsertGrade(grade *entities.Grade) (*entities.Grade, error) {
	grade, err := s.gradeRepo.InsertGrade(grade)
	if err != nil {
		return nil, services.HandleError(err, "Failed to insert grade", 500)
	}

	return grade, nil
}

func (s *teacherService) GetGradeByID(gradeID string) (*entities.Grade, error) {
	grade, err := s.gradeRepo.GetGradeByID(gradeID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get grade", 500)
	}

	return grade, nil
}

func (s *teacherService) GetAllGradeByStudentID(studentID string) ([]entities.Grade, error) {
	grades, err := s.gradeRepo.GetAllGradeByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get grades", 500)
	}

	return grades, nil
}

func (s *teacherService) GetAllGradeBySubjectID(subjectID string) ([]entities.Grade, error) {
	grades, err := s.gradeRepo.GetAllGradeBySubjectID(subjectID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get grades", 500)
	}

	return grades, nil
}

func (s *teacherService) GetAllGrade() ([]entities.Grade, error) {
	grades, err := s.gradeRepo.GetAllGrade()
	if err != nil {
		return nil, services.HandleError(err, "Failed to get grades", 500)
	}

	return grades, nil
}

func (s *teacherService) FilterBySemester(semester string) ([]entities.Grade, error) {
	grades, err := s.gradeRepo.FilterBySemester(semester)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get grades", 500)
	}

	return grades, nil
}

func (s *teacherService) FilterByAcademicYear(academicYear string) ([]entities.Grade, error) {
	grades, err := s.gradeRepo.FilterByAcademicYear(academicYear)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get grades", 500)
	}

	return grades, nil
}

func (s *teacherService) FilterBySemesterAndAcademicYear(semester, academicYear string) ([]entities.Grade, error) {
	grades, err := s.gradeRepo.FilterBySemesterAndAcademicYear(semester, academicYear)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get grades", 500)
	}
	return grades, nil
}

func (s *teacherService) UpdateGrade(grade *entities.Grade) (*entities.Grade, error) {
	grade, err := s.gradeRepo.UpdateGrade(grade)
	if err != nil {
		return nil, services.HandleError(err, "Failed to update grade", 500)
	}

	return grade, nil
}
