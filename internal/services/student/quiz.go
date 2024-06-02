package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentQuizService interface {
	GetQuiz(userID string) ([]entities.Quiz, error)
	GetQuizByID(id string) (*entities.Quiz, error)
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

func (s *studentService) GetQuizByID(id string) (*entities.Quiz, error) {
	quiz, err := s.quizRepo.GetQuiz(id)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch quiz", 500)
	}

	return quiz, nil
}
