package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherStudentAchivementService interface {
	GetAllAchivement() ([]entities.StudentAchivement, error)
	GetAchivementByID(achivementID string) (*entities.StudentAchivement, error)
	GetAllAchivementByStudentID(studentID string) ([]entities.StudentAchivement, error)
	UpdateAchivement(id string, achivement *entities.StudentAchivement) (*entities.StudentAchivement, error)
	DeleteAchivement(achivementID string) error
}

func (s *teacherService) GetAllAchivement() ([]entities.StudentAchivement, error) {
	achivements, err := s.achivementRepo.GetAllAchivement()
	if err != nil {
		return nil, services.HandleError(err, "Failed to get achivements", 500)
	}

	return achivements, nil
}

func (s *teacherService) GetAchivementByID(achivementID string) (*entities.StudentAchivement, error) {
	achivement, err := s.achivementRepo.GetAchivementByID(achivementID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get achivement", 500)
	}

	return achivement, nil
}

func (s *teacherService) GetAllAchivementByStudentID(studentID string) ([]entities.StudentAchivement, error) {
	achivements, err := s.achivementRepo.GetAllAchivementByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get achivements", 500)
	}

	return achivements, nil
}

func (s *teacherService) UpdateAchivement(id string, achivement *entities.StudentAchivement) (*entities.StudentAchivement, error) {
	achivement, err := s.achivementRepo.UpdateAchivement(id, achivement)
	if err != nil {
		return nil, services.HandleError(err, "Failed to update achivement", 500)
	}

	return achivement, nil
}

func (s *teacherService) DeleteAchivement(achivementID string) error {
	err := s.achivementRepo.DeleteAchivement(achivementID)
	if err != nil {
		return services.HandleError(err, "Failed to delete achivement", 500)
	}

	return nil
}
