package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentAssignmentService interface {
	SubmitAssignment(assignment *entities.StudentAssignment) error
	GetAssignment(taskID string) (*entities.StudentAssignment, error)
}

func (s *studentService) SubmitAssignment(assignment *entities.StudentAssignment) error {
	err := s.assignmentRepo.Insert(assignment)
	if err != nil {
		return services.HandleError(err, "Failed to submit assignment", 500)
	}

	return nil
}

func (s *studentService) GetAssignment(taskID string) (*entities.StudentAssignment, error) {
	assignment, err := s.assignmentRepo.FindByTaskID(taskID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch assignment", 500)
	}

	return assignment, nil
}
