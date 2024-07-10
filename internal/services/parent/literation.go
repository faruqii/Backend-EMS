package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type ParentStudentLiterationService interface {
	GetStudentLiterations(userID string) ([]entities.Literation, error)
	GetStudentLiterationDetail(literationID string) (*entities.Literation, error)
}

func (s *parentService) GetStudentLiterations(userID string) ([]entities.Literation, error) {
	parentID, err := s.tokenRepo.GetParentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get parent ID by user ID", 500)
	}

	studentID, err := s.parentRepo.GetStudentIDByParentID(parentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get student ID by parent ID", 500)
	}

	literations, err := s.literationRepo.GetLiterationByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get literation by student ID", 500)
	}

	return literations, nil

}

func (s *parentService) GetStudentLiterationDetail(literationID string) (*entities.Literation, error) {
	literation, err := s.literationRepo.GetLiterationByID(literationID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get literation by ID", 500)
	}

	return literation, nil
}
