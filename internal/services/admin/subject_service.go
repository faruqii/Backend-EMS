package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type AdminSubjectService interface {
	CreateSubject(subject *entities.Subject) error
	GetAllSubject() ([]entities.Subject, error)
	FindSubjectByID(id string) (*entities.Subject, error)
	IsTeacherAssignedToSubject(teacherID, subjectID string) (bool, error)
	GetClassesByPrefix(classPrefix string) ([]dto.ClassResponse, error)
	GetSubjectsByClassPrefix(classPrefix string) ([]dto.SubjectResponse, error)
	GetClassSubjectsByPrefixAndSubject(classPrefix string, subjectID string) ([]entities.ClassSubject, error)
	UpdateSubject(subjectID string, subject *entities.Subject) error
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

func (s *adminService) GetClassesByPrefix(classPrefix string) ([]dto.ClassResponse, error) {
	classes, err := s.classRepo.GetClassesByPrefix(classPrefix)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch classes", 500)
	}

	var classResponses []dto.ClassResponse
	for _, class := range classes {
		classResponses = append(classResponses, dto.ClassResponse{
			ID:              class.ID,
			Name:            class.Name,
			HomeRoomTeacher: class.HomeRoomTeacher.Name,
		})
	}
	return classResponses, nil
}

func (s *adminService) GetSubjectsByClassPrefix(classPrefix string) ([]dto.SubjectResponse, error) {
	subjects, err := s.subjectRepo.GetSubjectsByClassPrefix(classPrefix)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch subjects", 500)
	}

	var subjectResponses []dto.SubjectResponse
	for _, subject := range subjects {
		subjectResponses = append(subjectResponses, dto.SubjectResponse{
			ID:          subject.ID,
			Name:        subject.Name,
			Description: subject.Description,
			Semester:    subject.Semester,
		})
	}
	return subjectResponses, nil
}

// AdminService.go
func (s *adminService) GetClassSubjectsByPrefixAndSubject(classPrefix string, subjectID string) ([]entities.ClassSubject, error) {
	return s.subjectRepo.GetClassSubjectsByPrefixAndSubject(classPrefix, subjectID)
}

func (s *adminService) UpdateSubject(subjectID string, subject *entities.Subject) error {
	// find subject by id
	_, err := s.subjectRepo.FindByID(subjectID)
	if err != nil {
		return services.HandleError(err, "Failed to fetch subject", 500)
	}

	// update subject
	err = s.subjectRepo.Update(subjectID, subject)
	if err != nil {
		return services.HandleError(err, "Failed to update subject", 500)
	}

	return nil
}
