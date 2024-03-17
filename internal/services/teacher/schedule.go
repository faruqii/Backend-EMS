package services

import (
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherScheduleService interface {
	GetTodaySchedule(teacherID string, dayOfWeek time.Weekday) ([]entities.Schedule, error)
	GetTeacherIDByUserID(userID string) (string, error)
}

func (s *teacherService) GetTodaySchedule(teacherID string, dayOfWeek time.Weekday) ([]entities.Schedule, error) {
	schedules, err := s.scheduleRepo.GetTeacherTodaySchedule(teacherID, dayOfWeek)
	return schedules, services.HandleError(err, "Failed to fetch schedule", 500)
}

func (s *teacherService) GetTeacherIDByUserID(userID string) (string, error) {
	teacher, err := s.tokenRepo.GetTeacherIDByUserID(userID)
	return teacher, services.HandleError(err, "Failed to fetch teacher", 500)
}
