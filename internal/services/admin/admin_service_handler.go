package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/Magetan-Boyz/Backend/internal/services"
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

func (s *adminService) handleError(err error, message string, statusCode int) error {
	if err != nil {
		return &services.ErrorMessages{
			Message:    message,
			StatusCode: statusCode,
		}
	}
	return nil
}
