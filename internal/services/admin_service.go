package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
)

//go:generate mockgen -source=admin_service.go -destination=mock_admin_service.go -package=mock
type AdminService interface {
	CreateSubject(subject *entities.Subject) error
	CreateTeacher(teacher *entities.Teacher) error
}

type adminService struct {
	subjectRepository repositories.SubjectRepository
	teacherRepostory  repositories.TeacherRepository
}

func NewAdminService(subjectRepository repositories.SubjectRepository, teacherRepostory repositories.TeacherRepository) *adminService {
	return &adminService{
		subjectRepository: subjectRepository,
		teacherRepostory:  teacherRepostory,
	}
}

func (s *adminService) CreateSubject(subject *entities.Subject) error {
	err := s.subjectRepository.Create(subject)
	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to create subject",
			StatusCode: 500,
		}
	}
	return nil
}

func (s *adminService) CreateTeacher(teacher *entities.Teacher) error {
	err := s.teacherRepostory.Create(teacher)
	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to create teacher",
			StatusCode: 500,
		}
	}
	return nil
}
