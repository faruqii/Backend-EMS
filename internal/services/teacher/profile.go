package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherProfileService interface {
	GetMyProfile(userID string) (*entities.Teacher, error)
}

func (s *teacherService) GetMyProfile(userID string) (*entities.Teacher, error) {
	teacherID, err := s.tokenRepo.GetTeacherIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get teacher by user ID", 500)
	}

	teacher, err := s.teacherRepo.GetMyProfile(teacherID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get teacher profile", 500)
	}

	return teacher, nil
}
