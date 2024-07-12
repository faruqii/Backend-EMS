package service

import (
	"log"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentAssignmentService interface {
	SubmitAssignment(assignment *entities.StudentAssignment) error
	GetAssignment(taskID string) (*entities.StudentAssignment, error)
	SubmitQuiz(quizAssignment *entities.StudentQuizAssignment) error
	GetMyQuizAssignment(userID string, subjectID string) ([]entities.StudentQuizAssignment, error)
	UpdateTaskSubmission(assignmentID string, assignment *entities.StudentAssignment) error
}

func (s *studentService) SubmitAssignment(assignment *entities.StudentAssignment) error {
	err := s.assignmentRepo.Insert(assignment)
	if err != nil {
		return services.HandleError(err, "Failed to submit assignment", 500)
	}

	return nil
}

func (s *studentService) GetAssignment(taskID string) (*entities.StudentAssignment, error) {
	assignment, err := s.assignmentRepo.FindByTaskID(taskID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch assignment", 500)
	}

	return assignment, nil
}

func (s *studentService) SubmitQuiz(quizAssignment *entities.StudentQuizAssignment) error {
	// Get student id from token
	studentID, err := s.tokenRepo.GetStudentIDByUserID(quizAssignment.StudentID)
	if err != nil {
		return services.HandleError(err, "Failed to get student id", 500)
	}

	// Get quiz type by QuizID
	quizType, err := s.quizRepo.GetQuizType(quizAssignment.QuizID)
	if err != nil {
		return services.HandleError(err, "Failed to get quiz type", 500)
	}

	// Initialize default values
	quizAssignment.StudentID = studentID
	quizAssignment.SubmitAt = time.Now()
	quizAssignment.Status = "submitted"

	// Check quiz type
	if quizType == "Multiple Choice" {
		totalQuestions, err := s.quizRepo.CountQuestion(quizAssignment.QuizID)
		if err != nil {
			return services.HandleError(err, "Failed to count questions", 500)
		}
		log.Printf("Total Questions: %d\n", totalQuestions)

		correctAnswers, err := s.quizRepo.MatchAnswer(quizAssignment.QuizID, quizAssignment.Answers)
		if err != nil {
			return services.HandleError(err, "Failed to match answers", 500)
		}
		log.Printf("Correct Answers: %d\n", correctAnswers)

		grade := float64(correctAnswers) / float64(totalQuestions) * 100
		log.Printf("Calculated Grade: %f\n", grade)
		quizAssignment.Grade = grade

	} else if quizType == "Essay" {
		quizAssignment.Status = "waiting for graded"
		quizAssignment.Grade = 0
	}

	// Insert the assignment into the database
	if err := s.assignmentRepo.InsertQuiz(quizAssignment); err != nil {
		return services.HandleError(err, "Failed to submit quiz", 500)
	}

	return nil
}

func (s *studentService) GetMyQuizAssignment(userID string, subjectID string) ([]entities.StudentQuizAssignment, error) {
	// Get student id from token
	studentID, err := s.tokenRepo.GetStudentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get student id", 500)
	}

	// Get quiz assignment by student id
	quizAssignments, err := s.assignmentRepo.GetMyQuizAssignment(studentID, subjectID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get quiz assignment", 500)
	}

	return quizAssignments, nil
}

func (s *studentService) UpdateTaskSubmission(assignmentID string, assignment *entities.StudentAssignment) error {
	// check if the deadline has passed
	task, err := s.taskRepo.GetTask(assignment.TaskID)
	if err != nil {
		return services.HandleError(err, "Failed to get task", 500)
	}

	if time.Now().After(task.Deadline) {
		return services.HandleError(nil, "The deadline has passed", 400)
	}

	// Update the assignment
	if err := s.assignmentRepo.UpdateTaskSubmission(assignmentID, assignment); err != nil {
		return services.HandleError(err, "Failed to update assignment", 500)
	}

	return nil
}
