package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type ParentStudentAchievementService interface {
	GetStudentAchievement(userID string) ([]entities.StudentAchivement, error)
}

func (s *parentService) GetStudentAchievement(userID string) ([]entities.StudentAchivement, error) {
	parentID, err := s.tokenRepo.GetParentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch parent", 500)
	}

	studentID, err := s.parentRepo.GetStudentIDByParentID(parentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student", 500)
	}

	achivements, err := s.achivementRepo.GetAllAchivementByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student achivements", 500)
	}

	return achivements, nil
}
