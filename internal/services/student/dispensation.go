package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentDispensationService interface {
	CreateDispensation(userID string, dispensation *entities.Dispensation) (*entities.Dispensation, error)
	GetDispenpationByID(dispensationID string) (*entities.Dispensation, error)
	GetMyDispensations(userID string) ([]entities.Dispensation, error)
}

func (s *studentService) CreateDispensation(userID string, dispensation *entities.Dispensation) (*entities.Dispensation, error) {
	studentID, err := s.tokenRepo.GetStudentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get student id", 500)
	}

	dispensation.StudentID = studentID
	dispensation, err = s.dispensationRepo.InsertDispensation(dispensation)
	if err != nil {
		return nil, services.HandleError(err, "Failed to create dispensation", 500)
	}

	return dispensation, nil
}

func (s *studentService) GetDispenpationByID(dispensationID string) (*entities.Dispensation, error) {
	dispensation, err := s.dispensationRepo.GetDispensationByID(dispensationID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get dispensation", 500)
	}

	return dispensation, nil
}

func (s *studentService) GetMyDispensations(userID string) ([]entities.Dispensation, error) {
	studentID, err := s.tokenRepo.GetStudentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get student id", 500)
	}

	dispensations, err := s.dispensationRepo.GetDispensationsByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get dispensations", 500)
	}

	return dispensations, nil
}
