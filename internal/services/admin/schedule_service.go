package services

import "github.com/Magetan-Boyz/Backend/internal/domain/entities"

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
