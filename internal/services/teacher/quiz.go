package services

import (
	"fmt"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type QuizService interface {
	CreateQuiz(quiz *entities.Quiz, questions []entities.Question) error
	GetQuizByTeacherID(userID string) ([]entities.Quiz, error)
	GetAllQuizAssignment(quizID string) ([]entities.StudentQuizAssignment, error)
	GradeStudentQuiz(quizAssignmentID string, status string, grade float64) error
}

func (s *teacherService) CreateQuiz(quiz *entities.Quiz, questions []entities.Question) error {
	// Check if the teacher teaches the class
	isTeachingClass, err := s.classRepo.IsTeacherTeachTheClass(quiz.ClassID)
	if err != nil {
		return services.HandleError(err, "Failed to check if teacher teaches the class", 500)
	}

	if !isTeachingClass {
		return services.HandleError(fmt.Errorf("teacher does not teach this class"), "Teacher does not teach the class", 400)
	}

	// Check if the teacher is assigned to the subject
	isAssignedToSubject, err := s.subjectRepo.IsTeacherAssignedToSubject(quiz.TeacherID, quiz.SubjectID)
	if err != nil {
		return services.HandleError(err, "Failed to check if teacher teaches the subject", 500)
	}

	if !isAssignedToSubject {
		return services.HandleError(fmt.Errorf("teacher is not assigned to this subject"), "Teacher is not assigned to the subject", 400)
	}

	// Insert the quiz into the repository
	err = s.quizRepo.Insert(quiz)
	if err != nil {
		return services.HandleError(err, "Failed to create quiz", 500)
	}

	// Insert the questions into the repository
	for i := range questions {
		questions[i].QuizID = quiz.ID
	}

	err = s.quizRepo.CreateQuestion(questions)
	if err != nil {
		return services.HandleError(err, "Failed to create questions", 500)
	}

	return nil
}

func (s *teacherService) GetQuizByTeacherID(userID string) ([]entities.Quiz, error) {
	// get teacherID from token
	teacherID, err := s.tokenRepo.GetTeacherIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch teacher", 500)
	}

	quiz, err := s.quizRepo.GetQuizByTeacherID(teacherID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch quiz", 500)
	}

	return quiz, nil
}

func (s *teacherService) GetAllQuizAssignment(quizID string) ([]entities.StudentQuizAssignment, error) {
	quizAssignment, err := s.studentAssignmentRepo.GetAllQuizAssignment(quizID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch quiz assignment", 500)
	}

	return quizAssignment, nil
}

func (s *teacherService) GradeStudentQuiz(quizAssignmentID string, status string, grade float64) error {
	err := s.studentAssignmentRepo.GradeStudentQuiz(quizAssignmentID, status, grade)
	if err != nil {
		return services.HandleError(err, "Failed to grade student quiz", 500)
	}

	return nil
}
