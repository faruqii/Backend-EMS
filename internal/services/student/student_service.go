package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentService interface {
	GetScedule(userID string) ([]entities.Schedule, error)
	GetTask(userID string) (*entities.Task, error)
}

type studentService struct {
	scheduleRepo repositories.ScheduleRepository
	taskRepo     repositories.TaskRepository
	studentRepo  repositories.StudentRepository
	tokenRepo    repositories.TokenRepository
}

func NewStudentService(
	scheduleRepo repositories.ScheduleRepository,
	taskRepo repositories.TaskRepository,
	studentRepo repositories.StudentRepository,
	tokenRepo repositories.TokenRepository,
) *studentService {
	return &studentService{
		scheduleRepo: scheduleRepo,
		taskRepo:     taskRepo,
		studentRepo:  studentRepo,
		tokenRepo:    tokenRepo,
	}
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

func (s *studentService) GetTask(userID string) (*entities.Task, error) {
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
