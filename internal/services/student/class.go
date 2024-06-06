package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentClassService interface {
	MyClass(userID string) (*entities.Class, error)
	MySubjects(userID string) ([]entities.ClassSubject, error)
}

func (s *studentService) MyClass(userID string) (*entities.Class, error) {
	studentID, err := s.tokenRepo.GetStudentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student", 500)
	}

	classID, err := s.studentRepo.FindStudentClassIDByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student class", 500)
	}

	class, err := s.classRepo.FindByID(classID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch class", 500)
	}

	return class, nil
}

func (s *studentService) MySubjects(userID string) ([]entities.ClassSubject, error) {
	studentID, err := s.tokenRepo.GetStudentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student", 500)
	}

	classID, err := s.studentRepo.FindStudentClassIDByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student class", 500)
	}

	subjects, err := s.subjectRepo.GetClassSubjects(classID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch subjects", 500)
	}

	return subjects, nil
}
