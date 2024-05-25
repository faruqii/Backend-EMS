package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentAssignmentService interface {
	SubmitAssignment(assignment *entities.StudentAssignment) error
}

func (s *studentService) SubmitAssignment(assignment *entities.StudentAssignment) error {
	err := s.assignmentRepo.Insert(assignment)
	if err != nil {
		return services.HandleError(err, "Failed to submit assignment", 500)
	}

	return nil
}
