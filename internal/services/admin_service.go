package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	AdminSubjectService
	AdminTeacherService
	AdminClassService
	AdminScheduleService
}

type adminService struct {
	subjectRepo  repositories.SubjectRepository
	teacherRepo  repositories.TeacherRepository
	userRepo     repositories.UserRepository
	roleRepo     repositories.RoleRepository
	classRepo    repositories.ClassRepository
	scheduleRepo repositories.ScheduleRepository
}

func NewAdminService(
	subjectRepo repositories.SubjectRepository,
	teacherRepo repositories.TeacherRepository,
	userRepo repositories.UserRepository,
	roleRepo repositories.RoleRepository,
	classRepo repositories.ClassRepository,
	scheduleRepo repositories.ScheduleRepository,
) *adminService {
	return &adminService{
		subjectRepo:  subjectRepo,
		teacherRepo:  teacherRepo,
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		classRepo:    classRepo,
		scheduleRepo: scheduleRepo,
	}
}

// AdminSubjectService methods
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

// AdminTeacherService methods
type AdminTeacherService interface {
	CreateTeacher(teacher *entities.Teacher) error
	GetAllTeacher() ([]entities.Teacher, error)
	AssignTeacherToSubject(teacherID, SubjectID string) error
	FindTeacherByID(id string) (*entities.Teacher, error)
	GetTeachersBySubjectID(subjectID string) ([]dto.TeacherSubjectResponse, error)
	GetTeacherSubjects(teacherID string) ([]dto.TeacherSubjectsResponse, error)
	UpdateTeacherHomeroomStatus(teacherID string, isHomeroom bool) error
}

func (s *adminService) CreateTeacher(teacher *entities.Teacher) error {
	_, err := s.userRepo.FindByUsername(teacher.User.Username)
	if err == nil {
		return s.handleError(err, "Username already exist", 400)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(teacher.User.Password), bcrypt.MinCost)
	if err != nil {
		return s.handleError(err, "Failed to hash password", 500)
	}

	teacher.User.Password = string(hashedPassword)

	err = s.teacherRepo.Create(teacher)
	if err != nil {
		return s.handleError(err, "Failed to create teacher", 500)
	}

	err = s.roleRepo.AssignUserRole(teacher.User.ID, "teacher")
	return s.handleError(err, "Failed to assign role to teacher", 500)
}

func (s *adminService) GetAllTeacher() ([]entities.Teacher, error) {
	teachers, err := s.teacherRepo.GetAll()
	return teachers, s.handleError(err, "Failed to fetch teachers", 500)
}

func (s *adminService) AssignTeacherToSubject(teacherID, SubjectID string) error {
	isAssigned, err := s.subjectRepo.IsTeacherAssignedToSubject(teacherID, SubjectID)
	if err != nil {
		return s.handleError(err, "Failed to check if teacher is assigned to subject", 500)
	}
	if isAssigned {
		return s.handleError(err, "Teacher already assigned to subject", 400)
	}

	err = s.subjectRepo.AssignTeacherToSubject(teacherID, SubjectID)
	return s.handleError(err, "Failed to assign teacher to subject", 500)
}

func (s *adminService) FindTeacherByID(id string) (*entities.Teacher, error) {
	teacher, err := s.teacherRepo.FindByID(id)
	return teacher, s.handleError(err, "Failed to fetch teacher", 500)
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
		return nil, s.handleError(err, "Failed to fetch teacher subjects", 500)
	}

	var subjects []dto.TeacherSubjectsResponse

	if len(teacherSubjects) == 0 {
		subjects = append(subjects, dto.TeacherSubjectsResponse{
			TeacherName: "",
			SubjectName: []string{},
		})
	} else {
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

func (s *adminService) UpdateTeacherHomeroomStatus(teacherID string, isHomeroom bool) error {
	teacher, err := s.teacherRepo.FindByID(teacherID)
	if err != nil {
		return s.handleError(err, "Teacher not found", 400)
	}

	if teacher.IsHomeroom == isHomeroom {
		return s.handleError(err, "Homeroom status already updated", 400)
	}

	teacher.IsHomeroom = isHomeroom
	err = s.teacherRepo.Update(teacher)
	return s.handleError(err, "Failed to update teacher", 500)
}

// AdminClassService methods
type AdminClassService interface {
	CreateClass(class *entities.Class) error
	AssignHomeroomTeacher(classID, teacherID string) error
	FindClassByID(id string) (*entities.Class, error)
	GetAllClass() ([]entities.Class, error)
}

func (s *adminService) CreateClass(class *entities.Class) error {
	_, err := s.classRepo.FindByName(class.Name)
	if err == nil {
		return &ErrorMessages{
			Message:    "Class already exist",
			StatusCode: 400,
		}
	}

	err = s.classRepo.Insert(class)
	return s.handleError(err, "Failed to create class", 500)
}

func (s *adminService) AssignHomeroomTeacher(classID, teacherID string) error {
	teacher, err := s.teacherRepo.FindByID(teacherID)
	if err != nil {
		return s.handleError(err, "Teacher not found", 400)
	}

	class, err := s.classRepo.FindByID(classID)
	if err != nil {
		return s.handleError(err, "Class not found", 400)
	}

	if class.HomeRoomTeacherID != nil {
		return s.handleError(err, "Class already has a homeroom teacher", 400)
	}

	if teacher.IsHomeroom {
		return s.handleError(err, "Teacher is already a homeroom teacher", 400)
	}

	class.HomeRoomTeacherID = &teacherID
	if err := s.classRepo.Update(class); err != nil {
		return s.handleError(err, "Failed to assign teacher as homeroom", 500)
	}

	teacher.IsHomeroom = true
	if err := s.teacherRepo.Update(teacher); err != nil {
		return s.handleError(err, "Failed to update teacher", 500)
	}

	return nil
}

func (s *adminService) FindClassByID(id string) (*entities.Class, error) {
	class, err := s.classRepo.FindByID(id)
	return class, s.handleError(err, "Failed to fetch class", 500)
}

func (s *adminService) GetAllClass() ([]entities.Class, error) {
	classes, err := s.classRepo.GetAll()
	return classes, s.handleError(err, "Failed to fetch classes", 500)
}

// AdminScheduleService methods
type AdminScheduleService interface {
	CreateSchedule(schedule *entities.Schedule) error
	GetScheduleByID(id string) (*entities.Schedule, error)
}

func (s *adminService) CreateSchedule(schedule *entities.Schedule) error {
	err := s.scheduleRepo.Insert(schedule)
	return s.handleError(err, "Failed to create schedule", 500)
}

func (s *adminService) GetScheduleByID(id string) (*entities.Schedule, error) {
	schedule, err := s.scheduleRepo.GetScheduleByID(id)
	return schedule, s.handleError(err, "Failed to fetch schedule", 500)
}

// Helper method
func (s *adminService) handleError(err error, message string, statusCode int) error {
	if err != nil {
		return &ErrorMessages{
			Message:    message,
			StatusCode: statusCode,
		}
	}
	return nil
}
