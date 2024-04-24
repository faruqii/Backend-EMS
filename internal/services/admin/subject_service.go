package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type AdminSubjectService interface {
	CreateSubject(subject *entities.Subject) error
	GetAllSubject() ([]entities.Subject, error)
	FindSubjectByID(id string) (*entities.Subject, error)
	IsTeacherAssignedToSubject(teacherID, subjectID string) (bool, error)
}

func (s *adminService) CreateSubject(subject *entities.Subject) error {
	err := s.subjectRepo.Create(subject)
	return services.HandleError(err, "Failed to create subject", 500)
}

func (s *adminService) GetAllSubject() ([]entities.Subject, error) {
	subjects, err := s.subjectRepo.GetAll()
	return subjects, services.HandleError(err, "Failed to fetch subjects", 500)
}

func (s *adminService) FindSubjectByID(id string) (*entities.Subject, error) {
	subject, err := s.subjectRepo.FindByID(id)
	return subject, services.HandleError(err, "Failed to fetch subject", 500)
}

func (s *adminService) IsTeacherAssignedToSubject(teacherID, subjectID string) (bool, error) {
	isAssigned, err := s.subjectRepo.IsTeacherAssignedToSubject(teacherID, subjectID)
	return isAssigned, services.HandleError(err, "Failed to check teacher assignment", 500)
}
