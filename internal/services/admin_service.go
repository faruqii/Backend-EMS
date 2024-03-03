package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
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
	GetTeacherSubjects(teacherID string) ([]dto.TeacherSubjectsResponse, error)
	CreateClass(class *entities.Class) error
	AssignHomeroomTeacher(classID, teacherID string) error
	FindClassByID(id string) (*entities.Class, error)
	GetAllClass() ([]entities.Class, error)
	UpdateTeacherHomeroomStatus(teacherID string, status bool) error
	CreateSchedule(schedule *entities.Schedule) error
	GetScheduleByID(id string) (*entities.Schedule, error)
}

type adminService struct {
	subjectRepo  repositories.SubjectRepository
	teacherRepo  repositories.TeacherRepository
	userRepo     repositories.UserRepository
	roleRepo     repositories.RoleRepository
	classRepo    repositories.ClassRepository
	scheduleRepo repositories.ScheduleRepository
}

func NewAdminService(subjectRepo repositories.SubjectRepository, teacherRepo repositories.TeacherRepository, userRepo repositories.UserRepository, roleRepo repositories.RoleRepository, classRepo repositories.ClassRepository, scheduleRepo repositories.ScheduleRepository) *adminService {
	return &adminService{
		subjectRepo:  subjectRepo,
		teacherRepo:  teacherRepo,
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		classRepo:    classRepo,
		scheduleRepo: scheduleRepo,
	}
}

func (s *adminService) CreateSubject(subject *entities.Subject) error {
	err := s.subjectRepo.Create(subject)
	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to create subject",
			StatusCode: 500,
		}
	}
	return nil
}

func (s *adminService) GetAllSubject() ([]entities.Subject, error) {
	subjects, err := s.subjectRepo.GetAll()
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
	_, err := s.userRepo.FindByUsername(teacher.User.Username)
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

	err = s.teacherRepo.Create(teacher)
	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to create teacher",
			StatusCode: 500,
		}
	}

	// assign role to teacher
	err = s.roleRepo.AssignUserRole(teacher.User.ID, "teacher")
	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to assign role to teacher",
			StatusCode: 500,
		}
	}
	return nil
}

func (s *adminService) GetAllTeacher() ([]entities.Teacher, error) {
	teachers, err := s.teacherRepo.GetAll()
	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to fetch teachers",
			StatusCode: 500,
		}
	}
	return teachers, nil
}

func (s *adminService) AssignTeacherToSubject(teacherID, SubjectID string) error {
	isAssigned, err := s.subjectRepo.IsTeacherAssignedToSubject(teacherID, SubjectID)
	if err != nil {
		return err
	}
	if isAssigned {
		return &ErrorMessages{
			Message:    "Teacher already assigned to subject",
			StatusCode: 400,
		}
	}

	err = s.subjectRepo.AssignTeacherToSubject(teacherID, SubjectID)
	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to assign teacher to subject",
			StatusCode: 500,
		}
	}
	return nil
}

func (s *adminService) FindTeacherByID(id string) (*entities.Teacher, error) {
	teacher, err := s.teacherRepo.FindByID(id)
	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to fetch teacher",
			StatusCode: 500,
		}
	}
	return teacher, nil
}

func (s *adminService) FindSubjectByID(id string) (*entities.Subject, error) {
	subject, err := s.subjectRepo.FindByID(id)
	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to fetch subject",
			StatusCode: 500,
		}
	}
	return subject, nil
}

func (s *adminService) GetTeachersBySubjectID(subjectID string) ([]dto.TeacherSubjectResponse, error) {
	teacherSubjects, err := s.subjectRepo.GetTeachersBySubjectID(subjectID)
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

func (s *adminService) GetTeacherSubjects(teacherID string) ([]dto.TeacherSubjectsResponse, error) {
	teacherSubjects, err := s.subjectRepo.GetTeacherSubjects(teacherID)
	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to fetch subjects",
			StatusCode: 500,
		}
	}

	var subjects []dto.TeacherSubjectsResponse

	if len(teacherSubjects) == 0 {
		// Return an empty response with teacher_name as empty string
		subjects = append(subjects, dto.TeacherSubjectsResponse{
			TeacherName: "",
			SubjectName: []string{},
		})
	} else {
		// Construct response with teacher_name as the first teacher's name and subjects in a map
		subjectMap := make(map[string]bool)
		for _, ts := range teacherSubjects {
			subjectMap[ts.Subject.Name] = true
		}
		firstTeacherName := teacherSubjects[0].Teacher.Name
		subjects = append(subjects, dto.TeacherSubjectsResponse{
			TeacherName: firstTeacherName,
			SubjectName: make([]string, 0, len(subjectMap)),
		})
		for subject := range subjectMap {
			subjects[0].SubjectName = append(subjects[0].SubjectName, subject)
		}
	}

	return subjects, nil
}

func (s *adminService) CreateClass(class *entities.Class) error {

	// check if class is already exist by name
	_, err := s.classRepo.FindByName(class.Name)
	if err == nil {
		return &ErrorMessages{
			Message:    "Class already exist",
			StatusCode: 400,
		}
	}

	err = s.classRepo.Insert(class)
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
	teacher, err := s.teacherRepo.FindByID(teacherID)
	if err != nil {
		return &ErrorMessages{
			Message:    "Teacher not found",
			StatusCode: 400,
		}
	}

	// Check if class exists
	class, err := s.classRepo.FindByID(classID)
	if err != nil {
		return &ErrorMessages{
			Message:    "Class not found",
			StatusCode: 400,
		}
	}

	// Check if the class already has a homeroom teacher assigned
	if class.HomeRoomTeacherID != nil {
		return &ErrorMessages{
			Message:    "Class already has a homeroom teacher assigned",
			StatusCode: 400,
		}
	}

	// Check if the teacher is already designated as a homeroom teacher
	if teacher.IsHomeroom {
		return &ErrorMessages{
			Message:    "Teacher is already designated as a homeroom teacher",
			StatusCode: 400,
		}
	}

	// Update class with teacherID
	class.HomeRoomTeacherID = &teacherID
	if err := s.classRepo.Update(class); err != nil {
		return &ErrorMessages{
			Message:    "Failed to assign teacher as homeroom",
			StatusCode: 500,
		}
	}

	// Update teacher with isHomeroom
	teacher.IsHomeroom = true
	if err := s.teacherRepo.Update(teacher); err != nil {
		return &ErrorMessages{
			Message:    "Failed to update teacher",
			StatusCode: 500,
		}
	}

	return nil
}

func (s *adminService) FindClassByID(id string) (*entities.Class, error) {
	class, err := s.classRepo.FindByID(id)
	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to fetch class",
			StatusCode: 500,
		}
	}
	return class, nil
}

func (s *adminService) GetAllClass() ([]entities.Class, error) {
	classes, err := s.classRepo.GetAll()
	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to fetch classes",
			StatusCode: 500,
		}
	}
	return classes, nil
}

func (s *adminService) UpdateTeacherHomeroomStatus(teacherID string, status bool) error {
	teacher, err := s.teacherRepo.FindByID(teacherID)
	if err != nil {
		return &ErrorMessages{
			Message:    "Teacher not found",
			StatusCode: 400,
		}
	}

	teacher.IsHomeroom = status
	if err := s.teacherRepo.Update(teacher); err != nil {
		return &ErrorMessages{
			Message:    "Failed to update teacher",
			StatusCode: 500,
		}
	}
	return nil
}

func (s *adminService) CreateSchedule(schedule *entities.Schedule) error {
	err := s.scheduleRepo.Insert(schedule)
	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to create schedule",
			StatusCode: 500,
		}
	}
	return nil
}

func (s *adminService) GetScheduleByID(id string) (*entities.Schedule, error) {
	schedule, err := s.scheduleRepo.GetScheduleByID(id)
	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to fetch schedule",
			StatusCode: 500,
		}
	}
	return schedule, nil
}
