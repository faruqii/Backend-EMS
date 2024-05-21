package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
)

type TeacherService interface {
	TeacherScheduleService
	TeacherTaskService
}

type teacherService struct {
	teacherRepo  repositories.TeacherRepository
	scheduleRepo repositories.ScheduleRepository
	tokenRepo    repositories.TokenRepository
	taskRepo     repositories.TaskRepository
}

func NewTeacherService(
	teacherRepo repositories.TeacherRepository,
	scheduleRepo repositories.ScheduleRepository,
	tokenRepo repositories.TokenRepository,
	taskRepo repositories.TaskRepository,
) *teacherService {
	return &teacherService{
		teacherRepo:  teacherRepo,
		scheduleRepo: scheduleRepo,
		tokenRepo:    tokenRepo,
		taskRepo:     taskRepo,
	}
}
