package services

import (
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherService interface {
	GetTodaySchedule(teacherID string, dayOfWeek time.Weekday) ([]entities.Schedule, error)
	GetTeacherIDByUserID(userID string) (string, error)
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

func (s *teacherService) GetTodaySchedule(teacherID string, dayOfWeek time.Weekday) ([]entities.Schedule, error) {
	schedules, err := s.scheduleRepo.GetTeacherTodaySchedule(teacherID, dayOfWeek)
	return schedules, services.HandleError(err, "Failed to fetch schedule", 500)
}

func (s *teacherService) GetTeacherIDByUserID(userID string) (string, error) {
	teacher, err := s.tokenRepo.GetTeacherIDByUserID(userID)
	return teacher, services.HandleError(err, "Failed to fetch teacher", 500)
}
