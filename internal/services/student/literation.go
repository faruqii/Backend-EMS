package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentLiterationService interface {
	InsertLiteration(literation *entities.Literation) (*entities.Literation, error)
	GetLiterationByID(literationID string) (*entities.Literation, error)
	GetLiterationByStudentID(studentID string) ([]entities.Literation, error)
}

func (s *studentService) InsertLiteration(literation *entities.Literation) (*entities.Literation, error) {
	literation, err := s.literationRepo.InsertLiteration(literation)
	if err != nil {
		return nil, services.HandleError(err, "Failed to insert literation", 500)
	}

	return literation, nil
}

func (s *studentService) GetLiterationByID(literationID string) (*entities.Literation, error) {
	literation, err := s.literationRepo.GetLiterationByID(literationID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get literation by id", 500)
	}

	return literation, nil
}

func (s *studentService) GetLiterationByStudentID(studentID string) ([]entities.Literation, error) {
	literations, err := s.literationRepo.GetLiterationByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get literation by student id", 500)
	}

	return literations, nil
}
