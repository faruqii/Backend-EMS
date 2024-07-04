package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type ParentStudentDispensationService interface {
	GetStudentDispensations(userID string) ([]entities.Dispensation, error)
	GetStudentDispensationByID(dispensationID string) (*entities.Dispensation, error)
}

func (s *parentService) GetStudentDispensations(userID string) ([]entities.Dispensation, error) {
	parentID, err := s.tokenRepo.GetParentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get parent", 500)
	}

	studentID, err := s.parentRepo.GetStudentIDByParentID(parentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get student", 500)
	}

	dispensations, err := s.dispensationRepo.GetDispensationsByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get dispensations", 500)
	}

	return dispensations, nil
}

func (s *parentService) GetStudentDispensationByID(dispensationID string) (*entities.Dispensation, error) {
	dispensation, err := s.dispensationRepo.GetDispensationByID(dispensationID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get dispensation", 500)
	}

	return dispensation, nil
}
