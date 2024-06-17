package services

import (
	"errors"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherViolationService interface {
	CreateViolation(violation *entities.Violation) error
	GetAllViolation() ([]entities.Violation, error)
	GetViolationByID(id string) (*entities.Violation, error)
	GetViolationByStudentID(studentID string) ([]entities.Violation, error)
	UpdateViolation(violation *entities.Violation) error
	DeleteViolation(id string) error
}

func (s *teacherService) CreateViolation(violation *entities.Violation) error {
	err := s.violationRepo.Create(violation)
	if err != nil {
		return services.HandleError(errors.New("failed to create violation"), "Failed to create violation", 500)
	}
	return nil
}

func (s *teacherService) GetAllViolation() ([]entities.Violation, error) {
	violations, err := s.violationRepo.GetAll()
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch violations", 500)
	}
	return violations, nil
}

func (s *teacherService) GetViolationByID(id string) (*entities.Violation, error) {
	violation, err := s.violationRepo.GetByID(id)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch violation", 500)
	}
	return violation, nil
}

func (s *teacherService) GetViolationByStudentID(studentID string) ([]entities.Violation, error) {
	violations, err := s.violationRepo.GetByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch violations", 500)
	}
	return violations, nil
}

func (s *teacherService) UpdateViolation(violation *entities.Violation) error {
	err := s.violationRepo.Update(violation)
	if err != nil {
		return services.HandleError(err, "Failed to update violation", 500)
	}
	return nil
}

func (s *teacherService) DeleteViolation(id string) error {
	err := s.violationRepo.Delete(id)
	if err != nil {
		return services.HandleError(err, "Failed to delete violation", 500)
	}
	return nil
}
