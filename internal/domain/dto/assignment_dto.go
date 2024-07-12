package dto

import "time"

type StudentAssignmentRequest struct {
	Submission string `json:"submission"`
}

type StudentUpdateAssignmentRequest struct {
	Submission string `json:"submission"`
}

type StudentAssignmentResponse struct {
	ID         string    `json:"id"`
	Task       string    `json:"task"`
	Student    string    `json:"student"`
	Submission string    `json:"submission"`
	Grade      float64   `json:"grade"`
	Feedback   string    `json:"feedback"`
	SubmitAt   time.Time `json:"submit_at"`
}

type UpdateStudentTaskAssignmentRequest struct {
	Grade    float64 `json:"grade"`
	Feedback string  `json:"feedback"`
}

type SubmitQuizRequest struct {
	Answers []string `json:"answers"`
}

type StudentQuizAssignmentResponse struct {
	ID          string  `json:"id"`
	QuizID      string  `json:"quiz_id"`
	QuizName    string  `json:"quiz_name"`
	StudentName string  `json:"student_name"`
	NISN        string  `json:"nisn"`
	Grade       float64 `json:"grade"`
	Status      string  `json:"status"`
	SubmitAt    string  `json:"submit_at"`
}

type StudentQuizAssignmentAnswerResponse struct {
	StudentID   string               `json:"student_id"`
	StudentName string               `json:"student_name"`
	QuizID      string               `json:"quiz_id"`
	QuizTitle   string               `json:"quiz_title"`
	Questions   []QuestionWithAnswer `json:"questions"`
}

type QuestionWithAnswer struct {
	Question      string `json:"question"`
	Answer        string `json:"answer"`
	CorrectAnswer string `json:"correct_answer"`
}
