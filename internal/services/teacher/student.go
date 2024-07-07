package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherStudentService interface {
	GetAllStudentByClass(classID string) ([]entities.Student, error)
}

func (s *teacherService) GetAllStudentByClass(classID string) ([]entities.Student, error) {
	students, err := s.studentRepo.GetAllStudentsByClassID(classID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get students", 500)
	}

	return students, nil
}
