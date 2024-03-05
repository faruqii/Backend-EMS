package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

// AdminService is a contract for AdminService
// Use case: Admin can manage subjects, teachers, classes, and schedules
// Below is dependency injection for AdminService such as Subject, Teacher, Class, and Schedule Services
type AdminService interface {
	AdminSubjectService
	AdminTeacherService
	AdminClassService
	AdminScheduleService
}

// adminService is a struct for AdminService call repository layer so it can communicate with database
type adminService struct {
	subjectRepo  repositories.SubjectRepository
	teacherRepo  repositories.TeacherRepository
	userRepo     repositories.UserRepository
	roleRepo     repositories.RoleRepository
	classRepo    repositories.ClassRepository
	scheduleRepo repositories.ScheduleRepository
}

// NewAdminService is a constructor for adminService
// It will return adminService struct
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

// handleError is a function to handle error for adminService
func (s *adminService) handleError(err error, message string, statusCode int) error {
	if err != nil {
		return &services.ErrorMessages{
			Message:    message,
			StatusCode: statusCode,
		}
	}
	return nil
}
