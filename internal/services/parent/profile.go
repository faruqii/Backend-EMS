package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type ParentProfileService interface {
	GetMyProfile(userID string) (*entities.Parent, error)
}

func (s *parentService) GetMyProfile(userID string) (*entities.Parent, error) {
	parentID, err := s.tokenRepo.GetParentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get parent by user ID", 500)
	}

	parent, err := s.parentRepo.GetMyProfile(parentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get parent profile", 500)
	}

	return parent, nil
}
