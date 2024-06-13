package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentAchivementService interface {
	CreateAchivement(userID string, achivement *entities.StudentAchivement) (*entities.StudentAchivement, error)
	GetAchivementByID(achivementID string) (*entities.StudentAchivement, error)
	GetMyAchievements(userID string) ([]entities.StudentAchivement, error)
}

func (s *studentService) CreateAchivement(userID string, achivement *entities.StudentAchivement) (*entities.StudentAchivement, error) {
	studentID, err := s.tokenRepo.GetStudentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get student id", 500)
	}

	achivement.StudentID = studentID
	achivement, err = s.achivementRepo.InsertAchivement(achivement)
	if err != nil {
		return nil, services.HandleError(err, "Failed to create achivement", 500)
	}

	return achivement, nil
}

func (s *studentService) GetAchivementByID(achivementID string) (*entities.StudentAchivement, error) {
	achivement, err := s.achivementRepo.GetAchivementByID(achivementID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get achivement", 500)
	}

	return achivement, nil
}

func (s *studentService) GetMyAchievements(userID string) ([]entities.StudentAchivement, error) {
	studentID, err := s.tokenRepo.GetStudentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get student id", 500)
	}

	achivements, err := s.achivementRepo.GetAllAchivementByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get achivements", 500)
	}

	return achivements, nil
}