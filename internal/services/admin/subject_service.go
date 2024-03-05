package services

import "github.com/Magetan-Boyz/Backend/internal/domain/entities"

type AdminSubjectService interface {
	CreateSubject(subject *entities.Subject) error
	GetAllSubject() ([]entities.Subject, error)
	FindSubjectByID(id string) (*entities.Subject, error)
}

func (s *adminService) CreateSubject(subject *entities.Subject) error {
	err := s.subjectRepo.Create(subject)
	return s.handleError(err, "Failed to create subject", 500)
}

func (s *adminService) GetAllSubject() ([]entities.Subject, error) {
	subjects, err := s.subjectRepo.GetAll()
	return subjects, s.handleError(err, "Failed to fetch subjects", 500)
}

func (s *adminService) FindSubjectByID(id string) (*entities.Subject, error) {
	subject, err := s.subjectRepo.FindByID(id)
	return subject, s.handleError(err, "Failed to fetch subject", 500)
}
