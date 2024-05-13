package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
)

// AdminService is a contract for AdminService
// Use case: Admin can manage subjects, teachers, classes, and schedules
// Below is dependency injection for AdminService such as Subject, Teacher, Class, and Schedule Services
type AdminService interface {
	AdminSubjectService
	AdminTeacherService
	AdminClassService
	AdminScheduleService
	AdminStudentService
}

// adminService is a struct for AdminService call repository layer so it can communicate with database
type adminService struct {
	subjectRepo  repositories.SubjectRepository
	teacherRepo  repositories.TeacherRepository
	userRepo     repositories.UserRepository
	roleRepo     repositories.RoleRepository
	classRepo    repositories.ClassRepository
	scheduleRepo repositories.ScheduleRepository
	studentRepo  repositories.StudentRepository
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
	studentRepo repositories.StudentRepository,
) *adminService {
	return &adminService{
		subjectRepo:  subjectRepo,
		teacherRepo:  teacherRepo,
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		classRepo:    classRepo,
		scheduleRepo: scheduleRepo,
		studentRepo:  studentRepo,
	}
}
