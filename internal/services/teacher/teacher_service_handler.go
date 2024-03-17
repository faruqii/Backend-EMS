package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
)

type TeacherService interface {
	TeacherScheduleService
}

type teacherService struct {
	teacherRepo  repositories.TeacherRepository
	scheduleRepo repositories.ScheduleRepository
	tokenRepo    repositories.TokenRepository
}

func NewTeacherService(
	teacherRepo repositories.TeacherRepository,
	scheduleRepo repositories.ScheduleRepository,
	tokenRepo repositories.TokenRepository) *teacherService {
	return &teacherService{
		teacherRepo:  teacherRepo,
		scheduleRepo: scheduleRepo,
		tokenRepo:    tokenRepo,
	}
}
