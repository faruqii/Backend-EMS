package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type AttedanceService interface {
	CreateAttedance(attedance *entities.Atendance) (*entities.Atendance, error)
}

func (s *teacherService) CreateAttedance(attedance *entities.Atendance) (*entities.Atendance, error) {
	attedance, err := s.attedanceRepo.CreateAttedance(attedance)
	if err != nil {
		return nil, services.HandleError(err, "Failed to create attedance", 500)
	}
	return attedance, nil
}
