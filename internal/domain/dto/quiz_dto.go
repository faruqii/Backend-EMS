package dto

type CreateQuizRequest struct {
	TeacherID   string            `json:"teacher_id"`
	Title       string            `json:"title"`
	TypeOfQuiz  string            `json:"type_of_quiz"` // Quiz, UTS, UAS
	Description string            `json:"description"`
	Deadline    string            `json:"deadline"`
	Questions   []QuestionRequest `json:"questions"`
}

type QuizResponse struct {
	ID          string          `json:"id"`
	ClassID     string          `json:"class_id"`
	SubjectID   string          `json:"subject_id"`
	TeacherID   string          `json:"teacher_id"`
	Title       string          `json:"title"`
	TypeOfQuiz  string          `json:"type_of_quiz"` // Quiz, UTS, UAS
	Description string          `json:"description"`
	Deadline    string          `json:"deadline"`
	Questions   []QuestionBrief `json:"questions"`
}

type QuestionRequest struct {
	Text          string   `json:"text"`
	Options       []string `json:"options"`
	CorrectAnswer string   `json:"correct_answer"`
}

type QuestionBrief struct {
	Text    string   `json:"text"`
	Options []string `json:"options"`
}
