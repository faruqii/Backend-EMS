package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherTaskService interface {
	CreateTask(task *entities.Task) error
	GetTask(id string) (*entities.Task, error)
}

func (s *teacherService) CreateTask(task *entities.Task) error {
	err := s.taskRepo.Insert(task)
	return services.HandleError(err, "Failed to create task", 500)
}

func (s *teacherService) GetTask(id string) (*entities.Task, error) {
	task, err := s.taskRepo.GetTask(id)
	return task, services.HandleError(err, "Failed to fetch task", 500)
}
