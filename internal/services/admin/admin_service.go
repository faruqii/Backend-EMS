package services

import (
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/patrickmn/go-cache"
)

// AdminService is a contract for AdminService
// Use case: Admin can manage subjects, teachers, classes, and schedules
// Below is dependency injection for AdminService such as Subject, Teacher, Class, and Schedule Services
//
//go:generate mockgen -source=admin_service.go -destination=mock_admin_service.go -package=mock
type AdminService interface {
	AdminSubjectService
	AdminTeacherService
	AdminClassService
	AdminScheduleService
	AdminStudentService
	AdminParentService
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
	parentRepo   repositories.ParentRepository
	cache        *cache.Cache
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
	parentRepo repositories.ParentRepository,
) *adminService {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &adminService{
		subjectRepo:  subjectRepo,
		teacherRepo:  teacherRepo,
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		classRepo:    classRepo,
		scheduleRepo: scheduleRepo,
		studentRepo:  studentRepo,
		parentRepo:   parentRepo,
		cache:        c,
	}
}
