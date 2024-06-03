package dto

import "time"

type StudentAssignmentRequest struct {
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
