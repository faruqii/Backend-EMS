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
	// get student id from token
	studentID, err := s.tokenRepo.GetStudentIDByUserID(quizAssignment.StudentID)
	if err != nil {
		return services.HandleError(err, "Failed to get student id", 500)
	}

	totalQuestions, err := s.quizRepo.CountQuestion(quizAssignment.QuizID)
	log.Println(totalQuestions)
	if err != nil {
		return services.HandleError(err, "Failed to count questions", 500)
	}

	correctAnswers, err := s.quizRepo.MatchAnswer(quizAssignment.QuizID, quizAssignment.Answers)
	if err != nil {
		return services.HandleError(err, "Failed to match answers", 500)
	}

	grade := float64(correctAnswers) / float64(totalQuestions) * 100

	// Set submission details
	quizAssignment.StudentID = studentID
	quizAssignment.SubmitAt = time.Now()
	quizAssignment.Status = "submitted"
	quizAssignment.Grade = grade

	// Insert the assignment into the database
	if err := s.assignmentRepo.InsertQuiz(quizAssignment); err != nil {
		return services.HandleError(err, "Failed to submit quiz", 500)
	}

	return nil
}
