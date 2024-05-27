package dto

import "github.com/Magetan-Boyz/Backend/internal/domain/entities"

type CreateQuizRequest struct {
	TeacherID   string              `json:"teacher_id"`
	Title       string              `json:"title"`
	TypeOfQuiz  string              `json:"type_of_quiz"` // Quiz, UTS, UAS
	Description string              `json:"description"`
	Deadline    string              `json:"deadline"`
	Questions   []entities.Question `json:"questions"`
}

type CreateQuizResponse struct {
	ID          string              `json:"id"`
	ClassID     string              `json:"class_id"`
	SubjectID   string              `json:"subject_id"`
	TeacherID   string              `json:"teacher_id"`
	Title       string              `json:"title"`
	TypeOfQuiz  string              `json:"type_of_quiz"` // Quiz, UTS, UAS
	Description string              `json:"description"`
	Deadline    string              `json:"deadline"`
	Questions   []entities.Question `json:"questions"`
}

type QuestionRequest struct {
	Text          string   `json:"text"`
	Options       []string `json:"options"`
	CorrectAnswer string   `json:"correct_answer"`
}
