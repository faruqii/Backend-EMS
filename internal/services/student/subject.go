package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentSubjectRepository interface {
	GetSubjectsByClassID(classID string) ([]entities.ClassSubject, error)
	GetDetailSubject(subjectID string) (*entities.Subject, error)
	GetSubjectMatterBySubjectID(subjectID string) ([]entities.SubjectMattter, error)
	GetDetailSubjectMatter(subjectMatterID string) (*entities.SubjectMattter, error)
}

func (s *studentService) GetSubjectsByClassID(classID string) ([]entities.ClassSubject, error) {
	subjects, err := s.subjectRepo.GetAllSubjectInClass(classID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get subjects", 500)
	}

	return subjects, nil
}

func (s *studentService) GetDetailSubject(subjectID string) (*entities.Subject, error) {
	subject, err := s.subjectRepo.FindByID(subjectID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get subject", 500)
	}

	return subject, nil
}

func (s *studentService) GetSubjectMatterBySubjectID(subjectID string) ([]entities.SubjectMattter, error) {
	subjectMatters, err := s.subjectRepo.GetSubjectMatterBySubjectID(subjectID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get subject matters", 500)
	}

	return subjectMatters, nil
}

func (s *studentService) GetDetailSubjectMatter(subjectMatterID string) (*entities.SubjectMattter, error) {
	subjectMatter, err := s.subjectRepo.GetDetailSubjectMatter(subjectMatterID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get subject matter", 500)
	}

	return subjectMatter, nil
}
