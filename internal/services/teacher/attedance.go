package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type AttedanceService interface {
	CreateAttedance(attedance *entities.Atendance) (*entities.Atendance, error)
	GetAttedanceBySubjectID(subjectID string) ([]entities.Atendance, error)
	GetAttedanceByClassID(classID string) ([]entities.Atendance, error)
}

func (s *teacherService) CreateAttedance(attedance *entities.Atendance) (*entities.Atendance, error) {
	attedance, err := s.attedanceRepo.CreateAttedance(attedance)
	if err != nil {
		return nil, services.HandleError(err, "Failed to create attedance", 500)
	}
	return attedance, nil
}

func (s *teacherService) GetAttedanceBySubjectID(subjectID string) ([]entities.Atendance, error) {
	attedances, err := s.attedanceRepo.GetAttedanceBySubjectID(subjectID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch attedances", 500)
	}
	return attedances, nil
}

func (s *teacherService) GetAttedanceByClassID(classID string) ([]entities.Atendance, error) {
	attedances, err := s.attedanceRepo.GetAttedanceByClassID(classID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch attedances", 500)
	}
	return attedances, nil
}


