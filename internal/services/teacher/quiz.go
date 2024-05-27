package services

import (
	"fmt"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type QuizService interface {
	CreateQuiz(quiz *entities.Quiz, questions []entities.Question) error
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
