package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type ParentStudentAttedanceService interface {
	GetStudentAttedance(userID string) ([]entities.Atendance, error)
}

func (s *parentService) GetStudentAttedance(userID string) ([]entities.Atendance, error) {
	parentID, err := s.tokenRepo.GetParentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get parent", 500)
	}

	studentID, err := s.parentRepo.GetStudentIDByParentID(parentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get student", 500)
	}

	attedance, err := s.attedanceRepo.GetMyAttedance(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get attedance", 500)
	}

	return attedance, nil
}
