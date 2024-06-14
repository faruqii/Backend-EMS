package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type ParentStudentQuizService interface {
	GetQuizAssignment(userID string) ([]entities.StudentQuizAssignment, error)
}

func (s *parentService) GetQuizAssignment(userID string) ([]entities.StudentQuizAssignment, error) {
	parentID, err := s.tokenRepo.GetParentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch parent", 500)
	}

	studentID, err := s.parentRepo.GetStudentIDByParentID(parentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student", 500)
	}

	quiz, err := s.assignmentRepo.GetQuizByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch quiz", 500)
	}

	return quiz, nil
}
