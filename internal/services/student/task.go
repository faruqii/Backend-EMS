package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentTaskService interface {
	GetScedule(userID string) ([]entities.Schedule, error)
	GetTask(userID string) ([]entities.Task, error)
	GetStudentIDByUserID(userID string) (string, error)
}

func (s *studentService) GetScedule(userID string) ([]entities.Schedule, error) {
	studentID, err := s.tokenRepo.GetStudentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student", 500)
	}

	classID, err := s.studentRepo.FindStudentClassIDByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student class", 500)
	}

	schedules, err := s.scheduleRepo.FindByClassID(classID)
	return schedules, services.HandleError(err, "Failed to fetch schedules", 500)
}

func (s *studentService) GetTask(userID string) ([]entities.Task, error) {
	studentID, err := s.tokenRepo.GetStudentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student", 500)
	}

	classID, err := s.studentRepo.FindStudentClassIDByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch class", 500)
	}

	task, err := s.taskRepo.GetTaskByClassID(classID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch task", 500)
	}

	return task, nil
}

func (s *studentService) GetStudentIDByUserID(userID string) (string, error) {
	studentID, err := s.tokenRepo.GetStudentIDByUserID(userID)
	return studentID, services.HandleError(err, "Failed to fetch student", 500)
}
