package services

import "github.com/Magetan-Boyz/Backend/internal/domain/entities"

type AdminScheduleService interface {
	CreateSchedule(schedule *entities.Schedule) error
	GetScheduleByID(id string) (*entities.Schedule, error)
	GetPreloadSchedule() (*entities.Schedule, error)
	IsScheduleExists(classID, subjectID string) (bool, error)
}

func (s *adminService) CreateSchedule(schedule *entities.Schedule) error {
	err := s.scheduleRepo.Insert(schedule)
	return s.handleError(err, "Failed to create schedule", 500)
}

func (s *adminService) GetScheduleByID(id string) (*entities.Schedule, error) {
	schedule, err := s.scheduleRepo.GetScheduleByID(id)
	return schedule, s.handleError(err, "Failed to fetch schedule", 500)
}

func (s *adminService) GetPreloadSchedule() (*entities.Schedule, error) {
	schedules, err := s.scheduleRepo.GetPreloadSchedule()
	return schedules, s.handleError(err, "Failed to fetch schedules", 500)
}

func (s *adminService) IsScheduleExists(classID, subjectID string) (bool, error) {
	exists, err := s.scheduleRepo.IsScheduleExists(classID, subjectID)
	return exists, s.handleError(err, "Failed to check if schedule exists", 500)
}
