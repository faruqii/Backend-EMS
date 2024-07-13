package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type ParentTaskService interface {
	GetTask(userID string) ([]entities.Task, error)
	GetStudentAssignment(userID string) ([]entities.StudentAssignment, error)
}

func (s *parentService) GetTask(userID string) ([]entities.Task, error) {
	parentID, err := s.tokenRepo.GetParentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch parent", 500)
	}

	studentID, err := s.parentRepo.GetStudentIDByParentID(parentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student", 500)
	}

	classID, err := s.studentRepo.FindStudentClassIDByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student class", 500)
	}

	task, err := s.taskRepo.GetTaskByClassID(classID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch task", 500)
	}

	return task, nil
}

func (s *parentService) GetStudentAssignment(userID string) ([]entities.StudentAssignment, error) {
	parentID, err := s.tokenRepo.GetParentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch parent", 500)
	}

	studentID, err := s.parentRepo.GetStudentIDByParentID(parentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student", 500)
	}

	studentAssignment, err := s.assignmentRepo.GetStudentTaskSubmissions(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student assignment", 500)
	}

	return studentAssignment, nil
}
