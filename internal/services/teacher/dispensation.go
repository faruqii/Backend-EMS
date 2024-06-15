package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherDispensationService interface {
	GetDispensationByID(id string) (*entities.Dispensation, error)
	GetAllDispensations() ([]entities.Dispensation, error)
	GetDispensationsByStudentID(studentID string) ([]entities.Dispensation, error)
	UpdateDispensationStatus(dispensationID string, status string) (*entities.Dispensation, error)
}

func (s *teacherService) GetDispensationByID(id string) (*entities.Dispensation, error) {
	dispensation, err := s.dispensationRepo.GetDispensationByID(id)
	if err != nil {
		return nil, services.HandleError(err, "Dispensation not found", 404)
	}
	return dispensation, nil
}

func (s *teacherService) GetAllDispensations() ([]entities.Dispensation, error) {
	dispensations, err := s.dispensationRepo.GetAllDispensations()
	if err != nil {
		return nil, services.HandleError(err, "Dispensations not found", 404)
	}
	return dispensations, nil
}

func (s *teacherService) GetDispensationsByStudentID(studentID string) ([]entities.Dispensation, error) {
	dispensations, err := s.dispensationRepo.GetDispensationsByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Dispensations not found", 404)
	}
	return dispensations, nil
}

func (s *teacherService) UpdateDispensationStatus(dispensationID string, status string) (*entities.Dispensation, error) {
	dispensation, err := s.dispensationRepo.UpdateDispensationStatus(dispensationID, status)
	if err != nil {
		return nil, services.HandleError(err, "Failed to update dispensation status", 500)
	}
	return dispensation, nil
}
