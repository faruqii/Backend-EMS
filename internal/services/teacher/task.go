package services

import (
	"fmt"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherTaskService interface {
	CreateTask(task *entities.Task) error
	GetTask(id string) (*entities.Task, error)
	GetAllTasks(userID string) ([]entities.Task, error)
	GetStudentTaskAssignment(taskID string) ([]entities.StudentAssignment, error)
	UpdateStudentTaskAssignment(assignmentID string, grade float64, feedback string) error
	GetStudentTaskAssignmentDetail(assignmentID string) (*entities.StudentAssignment, error)
	UpdateTask(taskID string, task *entities.Task) error
	DeleteTask(taskID string) error
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

func (s *teacherService) GetAllTasks(userID string) ([]entities.Task, error) {
	teacherID, err := s.tokenRepo.GetTeacherIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch teacher", 500)
	}

	tasks, err := s.taskRepo.GetTaskByTeacherID(teacherID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch tasks", 500)
	}

	return tasks, nil
}

func (s *teacherService) GetStudentTaskAssignment(taskID string) ([]entities.StudentAssignment, error) {
	studentAssignments, err := s.studentAssignmentRepo.FindAll(taskID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student assignments", 500)
	}

	return studentAssignments, nil
}

func (s *teacherService) UpdateStudentTaskAssignment(assignmentID string, grade float64, feedback string) error {
	assignment, err := s.studentAssignmentRepo.FindByID(assignmentID)
	if err != nil {
		return services.HandleError(err, "Failed to fetch student assignment", 500)
	}

	assignment.Grade = grade
	assignment.Feedback = feedback

	err = s.studentAssignmentRepo.Update(assignment)
	if err != nil {
		return services.HandleError(err, "Failed to update student assignment", 500)
	}

	return nil
}

func (s *teacherService) GetStudentTaskAssignmentDetail(assignmentID string) (*entities.StudentAssignment, error) {
	assignment, err := s.studentAssignmentRepo.FindByID(assignmentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student assignment", 500)
	}

	return assignment, nil
}

func (s *teacherService) UpdateTask(taskID string, task *entities.Task) error {
	err := s.taskRepo.Update(taskID, task)
	return services.HandleError(err, "Failed to update task", 500)
}

func (s *teacherService) DeleteTask(taskID string) error {
	err := s.taskRepo.Delete(taskID)
	return services.HandleError(err, "Failed to delete task", 500)
}
