package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentViolationService interface {
	GetMyViolation(userID string) ([]entities.Violation, error)
	GetViolationByID(violationID string) (*entities.Violation, error)
}

func (s *studentService) GetMyViolation(userID string) ([]entities.Violation, error) {
	studentID, err := s.tokenRepo.GetStudentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Student not Found", 500)
	}

	violations, err := s.violationRepo.GetByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch violations", 500)
	}

	return violations, nil
}

func (s *studentService) GetViolationByID(violationID string) (*entities.Violation, error) {
	violation, err := s.violationRepo.GetByID(violationID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch violations", 500)
	}

	return violation, nil
}
