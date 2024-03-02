package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -source=admin_service.go -destination=mock_admin_service.go -package=mock
type AdminService interface {
	CreateSubject(subject *entities.Subject) error
	GetAllSubject() ([]entities.Subject, error)
	CreateTeacher(teacher *entities.Teacher) error
	GetAllTeacher() ([]entities.Teacher, error)
	AssignTeacherToSubject(teacherID, SubjectID string) error
	FindTeacherByID(id string) (*entities.Teacher, error)
	FindSubjectByID(id string) (*entities.Subject, error)
	GetTeachersBySubjectID(subjectID string) ([]dto.TeacherSubjectResponse, error)
	CreateClass(class *entities.Class) error
	AssignHomeroomTeacher(classID, teacherID string) error
	FindClassByID(id string) (*entities.Class, error)
}

type adminService struct {
	subjectRepository repositories.SubjectRepository
	teacherRepository repositories.TeacherRepository
	userRepository    repositories.UserRepository
	roleRepository    repositories.RoleRepository
	classRepository  repositories.ClassRepository
}

func NewAdminService(subjectRepository repositories.SubjectRepository, teacherRepository repositories.TeacherRepository, userRepository repositories.UserRepository, roleRepository repositories.RoleRepository, classRepository repositories.ClassRepository) *adminService {
	return &adminService{
		subjectRepository: subjectRepository,
		teacherRepository: teacherRepository,
		userRepository:    userRepository,
		roleRepository:    roleRepository,
		classRepository:  classRepository,
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

func (s *adminService) GetAllSubject() ([]entities.Subject, error) {
	subjects, err := s.subjectRepository.GetAll()
	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to fetch subjects",
			StatusCode: 500,
		}
	}
	return subjects, nil
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

func (s *adminService) AssignTeacherToSubject(teacherID, SubjectID string) error {
	isAssigned, err := s.subjectRepository.IsTeacherAssignedToSubject(teacherID, SubjectID)
	if err != nil {
		return err
	}
	if isAssigned {
		return &ErrorMessages{
			Message:    "Teacher already assigned to subject",
			StatusCode: 400,
		}
	}

	err = s.subjectRepository.AssignTeacherToSubject(teacherID, SubjectID)
	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to assign teacher to subject",
			StatusCode: 500,
		}
	}
	return nil
}

func (s *adminService) FindTeacherByID(id string) (*entities.Teacher, error) {
	teacher, err := s.teacherRepository.FindByID(id)
	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to fetch teacher",
			StatusCode: 500,
		}
	}
	return teacher, nil
}

func (s *adminService) FindSubjectByID(id string) (*entities.Subject, error) {
	subject, err := s.subjectRepository.FindByID(id)
	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to fetch subject",
			StatusCode: 500,
		}
	}
	return subject, nil
}

func (s *adminService) GetTeachersBySubjectID(subjectID string) ([]dto.TeacherSubjectResponse, error) {
	teacherSubjects, err := s.subjectRepository.GetTeachersBySubjectID(subjectID)
	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to fetch teachers",
			StatusCode: 500,
		}
	}

	var teachers []dto.TeacherSubjectResponse
	for _, ts := range teacherSubjects {
		teachers = append(teachers, dto.TeacherSubjectResponse{
			TeacherName: ts.Teacher.Name,
			SubjectName: ts.Subject.Name,
		})
	}
	return teachers, nil

}

func (s *adminService) CreateClass(class *entities.Class) error {
	err := s.classRepository.Insert(class)
	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to create class",
			StatusCode: 500,
		}
	}
	return nil
}

func (s *adminService) AssignHomeroomTeacher(classID, teacherID string) error {
	// Check if teacher exists
	teacher, err := s.teacherRepository.FindByID(teacherID)
	if err != nil {
		return &ErrorMessages{
			Message:    "Teacher not found",
			StatusCode: 400,
		}
	}

	// Check if class exists
	class, err := s.classRepository.FindByID(classID)
	if err != nil {
		return &ErrorMessages{
			Message:    "Class not found",
			StatusCode: 400,
		}
	}

	// Update class with teacherID
	class.HomeRoomTeacherID = &teacherID
	if err := s.classRepository.Update(class); err != nil {
		return &ErrorMessages{
			Message:    "Failed to assign teacher as homeroom",
			StatusCode: 500,
		}
	}

	// Update teacher with isHomeroom
	teacher.IsHomeroom = true
	if err := s.teacherRepository.Update(teacher); err != nil {
		return &ErrorMessages{
			Message:    "Failed to update teacher",
			StatusCode: 500,
		}
	}

	return nil
}

func (s *adminService) FindClassByID(id string) (*entities.Class, error) {
	class, err := s.classRepository.FindByID(id)
	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to fetch class",
			StatusCode: 500,
		}
	}
	return class, nil
}




