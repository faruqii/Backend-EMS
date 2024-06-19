package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type ParentStudentViolationService interface {
	GetStudentViolation(userID string) ([]entities.Violation, error)
	GetViolationByID(violationID string) (*entities.Violation, error)
}

func (s *parentService) GetStudentViolation(userID string) ([]entities.Violation, error) {
	// get parent by user id
	parentID, err := s.tokenRepo.GetParentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Parent not Found", 500)
	}

	// get student by parent id
	studentID, err := s.parentRepo.GetStudentIDByParentID(parentID)
	if err != nil {
		return nil, services.HandleError(err, "Student not Found", 500)
	}

	// get violations by student id
	violations, err := s.violationRepo.GetByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch violations", 500)
	}

	return violations, nil

}

func (s *parentService) GetViolationByID(violationID string) (*entities.Violation, error) {
	// get violation by id
	violation, err := s.violationRepo.GetByID(violationID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch violations", 500)
	}

	return violation, nil
}
