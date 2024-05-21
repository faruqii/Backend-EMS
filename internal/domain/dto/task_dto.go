package dto

import "time"

type TaskRequest struct {
	ClassID     string `json:"class_id"`
	SubjectID   string `json:"subject_id"`
	Title       string `json:"title"`
	TypeOfTask  string `json:"type_of_task"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Link        string `json:"link"`
}

type TaskResponse struct {
	ID          string    `json:"id"`
	ClassName   string    `json:"class"`
	SubjectName string    `json:"subject"`
	TeacherName string    `json:"teacher"`
	Title       string    `json:"title"`
	TypeOfTask  string    `json:"type_of_task"`
	Description string    `json:"description"`
	Deadline    string    `json:"deadline"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
