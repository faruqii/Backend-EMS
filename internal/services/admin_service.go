package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -source=admin_service.go -destination=mock_admin_service.go -package=mock
type AdminService interface {
	CreateSubject(subject *entities.Subject) error
	CreateTeacher(teacher *entities.Teacher) error
	GetAllTeacher() ([]entities.Teacher, error)
}

type adminService struct {
	subjectRepository repositories.SubjectRepository
	teacherRepository repositories.TeacherRepository
	userRepository    repositories.UserRepository
	roleRepository    repositories.RoleRepository
}

func NewAdminService(subjectRepository repositories.SubjectRepository, teacherRepository repositories.TeacherRepository, userRepository repositories.UserRepository, roleRepository repositories.RoleRepository) *adminService {
	return &adminService{
		subjectRepository: subjectRepository,
		teacherRepository: teacherRepository,
		userRepository:    userRepository,
		roleRepository:    roleRepository,
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

	// check if teacher is exist

	_, err := s.userRepository.FindByUsername(teacher.User.Username)
	if err == nil {
		return &ErrorMessages{
			Message:    "Teacher already exist",
			StatusCode: 400,
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(teacher.User.Password), bcrypt.MinCost)
	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to hash password",
			StatusCode: 500,
		}
	}

	teacher.User.Password = string(hashedPassword)

	err = s.teacherRepository.Create(teacher)
	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to create teacher",
			StatusCode: 500,
		}
	}

	// assign role to teacher
	err = s.roleRepository.AssignUserRole(teacher.User.ID, "teacher")
	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to assign role to teacher",
			StatusCode: 500,
		}
	}
	return nil
}

func (s *adminService) GetAllTeacher() ([]entities.Teacher, error) {
	teachers, err := s.teacherRepository.GetAll()
	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to fetch teachers",
			StatusCode: 500,
		}
	}
	return teachers, nil
}
