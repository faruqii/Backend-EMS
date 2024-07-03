package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type ParentStudentService interface {
	GetMyStudent(userID string) ([]entities.Student, error)
	GetStudentByID(studentID string) (*entities.Student, error)
}

func (s *parentService) GetMyStudent(userID string) ([]entities.Student, error) {
	parentID, err := s.tokenRepo.GetParentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get parent id", 500)
	}

	students, err := s.parentRepo.GetMyStudents(parentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch students", 500)
	}

	var studentResponses []entities.Student
	for _, student := range students {
		student, err := s.studentRepo.FindById(student.StudentID)
		if err != nil {
			return nil, services.HandleError(err, "Failed to fetch student", 500)
		}
		studentResponses = append(studentResponses, *student)
	}

	return studentResponses, nil
}

func (s *parentService) GetStudentByID(studentID string) (*entities.Student, error) {
	student, err := s.studentRepo.FindById(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student", 500)
	}

	return student, nil
}
