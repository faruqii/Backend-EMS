package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentQuizService interface {
	GetQuiz(userID string) ([]entities.Quiz, error)
	GetQuizQuestions(quizID string, page, pageSize int) ([]entities.Question, error)
	CountQuizQuestions(quizID string) (int, error)
	GetQuizByID(id string) (*entities.Quiz, error)
	GetMyQuizGrade(quizID, userID string) (*entities.StudentQuizAssignment, error)
	GetMyQuizSubmission(quizAssignmentID string) (*entities.StudentQuizAssignment, error)
}

func (s *studentService) GetQuiz(userID string) ([]entities.Quiz, error) {
	studentID, err := s.tokenRepo.GetStudentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student", 500)
	}

	classID, err := s.studentRepo.FindStudentClassIDByStudentID(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch class", 500)
	}

	quiz, err := s.quizRepo.GetQuizByClassID(classID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch quiz", 500)
	}

	return quiz, nil
}

func (s *studentService) GetQuizQuestions(quizID string, page, pageSize int) ([]entities.Question, error) {
	questions, err := s.quizRepo.GetQuestion(quizID, page, pageSize)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch questions", 500)
	}

	return questions, nil
}

func (s *studentService) CountQuizQuestions(quizID string) (int, error) {
	return s.quizRepo.CountQuestions(quizID)
}

func (s *studentService) GetQuizByID(id string) (*entities.Quiz, error) {
	quiz, err := s.quizRepo.GetQuiz(id)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch quiz", 500)
	}

	return quiz, nil
}

func (s *studentService) GetMyQuizGrade(quizID, userID string) (*entities.StudentQuizAssignment, error) {
	studentID, err := s.tokenRepo.GetStudentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student", 500)
	}

	quizAssignment, err := s.assignmentRepo.GetStudentQuizAssignment(quizID, studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch quiz assignment", 500)
	}

	return quizAssignment, nil
}

func (s *studentService) GetMyQuizSubmission(quizAssignmentID string) (*entities.StudentQuizAssignment, error) {
	quizAssignment, err := s.assignmentRepo.GetQuizAssignment(quizAssignmentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch quiz assignment", 500)
	}

	return quizAssignment, nil
}
