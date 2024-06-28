package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type AdminScheduleService interface {
	CreateSchedule(schedule *entities.Schedule) error
	GetScheduleByID(id string) (*entities.Schedule, error)
	GetPreloadSchedule() (*entities.Schedule, error)
	IsScheduleExists(classID, subjectID string) (bool, error)
	GetAllSchedule() ([]entities.Schedule, error)
	UpdateSchedule(schedule *entities.Schedule) error
	DeleteSchedule(id string) error
}

func (s *adminService) CreateSchedule(schedule *entities.Schedule) error {
	err := s.scheduleRepo.Insert(schedule)
	return services.HandleError(err, "Failed to create schedule", 500)
}

func (s *adminService) GetScheduleByID(id string) (*entities.Schedule, error) {
	schedule, err := s.scheduleRepo.GetScheduleByID(id)
	return schedule, services.HandleError(err, "Failed to fetch schedule", 500)
}

func (s *adminService) GetPreloadSchedule() (*entities.Schedule, error) {
	schedules, err := s.scheduleRepo.GetPreloadSchedule()
	return schedules, services.HandleError(err, "Failed to fetch schedules", 500)
}

func (s *adminService) IsScheduleExists(classID, subjectID string) (bool, error) {
	exists, err := s.scheduleRepo.IsScheduleExists(classID, subjectID)
	return exists, services.HandleError(err, "Failed to check if schedule exists", 500)
}

func (s *adminService) GetAllSchedule() ([]entities.Schedule, error) {
	schedules, err := s.scheduleRepo.GetAll()
	return schedules, services.HandleError(err, "Failed to fetch schedules", 500)
}

func (s *adminService) UpdateSchedule(schedule *entities.Schedule) error {
	err := s.scheduleRepo.Update(schedule)
	return services.HandleError(err, "Failed to update schedule", 500)
}

func (s *adminService) DeleteSchedule(id string) error {
	err := s.scheduleRepo.Delete(id)
	return services.HandleError(err, "Failed to delete schedule", 500)
}
