package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherStudentLiterationService interface {
	GetAllLiterations() ([]entities.Literation, error)
	FilterByClassID(classID string) ([]entities.Literation, error)
	UpdateLiterationFeedback(literationID string, feedback string) (*entities.Literation, error)
	GetLiterationByID(literationID string) (*entities.Literation, error)
	GetLiterationByStudentID(studentID string) ([]entities.Literation, error)
}

func (s *teacherService) GetAllLiterations() ([]entities.Literation, error) {
	literations, err := s.literationRepo.GetAllLiterations()
	if err != nil {
		return nil, services.HandleError(err, "Failed to get all literations", 500)
	}

	return literations, nil
}

func (s *teacherService) FilterByClassID(classID string) ([]entities.Literation, error) {
	literations, err := s.literationRepo.FilterByStudentClass(classID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to filter literation by class id", 500)
	}

	return literations, nil
}

func (s *teacherService) UpdateLiterationFeedback(literationID string, feedback string) (*entities.Literation, error) {
	literation, err := s.literationRepo.UpdateLiterationFeedback(literationID, feedback)
	if err != nil {
		return nil, services.HandleError(err, "Failed to update literation feedback", 500)
	}

	return literation, nil
}

func (s *teacherService) GetLiterationByID(literationID string) (*entities.Literation, error) {
	literation, err := s.literationRepo.GetLiterationByID(literationID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get literation by id", 500)
	}

	return literation, nil
}

func (s *teacherService) GetLiterationByStudentID(studentID string) ([]entities.Literation, error) {
	literations, err := s.literationRepo.GetLiterationByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get literation by student id", 500)
	}

	return literations, nil
}
