package services

import (
	"fmt"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherTaskService interface {
	CreateTask(task *entities.Task) error
	GetTask(id string) (*entities.Task, error)
}

func (s *teacherService) CreateTask(task *entities.Task) error {
	// Check if the teacher teaches the class
	isTeachingClass, err := s.classRepo.IsTeacherTeachTheClass(task.ClassID)
	if err != nil {
		return services.HandleError(err, "Failed to check if teacher teaches the class", 500)
	}
	if !isTeachingClass {
		return services.HandleError(fmt.Errorf("teacher does not teach this class"), "Teacher does not teach the class", 400)
	}

	// Check if the teacher is assigned to the subject
	isAssignedToSubject, err := s.subjectRepo.IsTeacherAssignedToSubject(task.TeacherID, task.SubjectID)
	if err != nil {
		return services.HandleError(err, "Failed to check if teacher teaches the subject", 500)
	}
	if !isAssignedToSubject {
		return services.HandleError(fmt.Errorf("teacher is not assigned to this subject"), "Teacher is not assigned to the subject", 400)
	}

	// Insert the task into the repository
	err = s.taskRepo.Insert(task)
	if err != nil {
		return services.HandleError(err, "Failed to create task", 500)
	}

	return nil
}

func (s *teacherService) GetTask(id string) (*entities.Task, error) {
	task, err := s.taskRepo.GetTask(id)
	return task, services.HandleError(err, "Failed to fetch task", 500)
}
