package dto

import "time"

type SubjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Semester    string `json:"semester"`
}

type SubjectResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Semester    string `json:"semester"`
}

type TeacherSubjectRequest struct {
	TeacherID []string `json:"teacher_id"`
}

type TeacherSubjectResponse struct {
	SubjectID   string `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	TeacherID   string `json:"teacher_id"`
	TeacherName string `json:"teacher_name"`
}

type TeacherSubjectsResponse struct {
	TeacherID   string            `json:"teacher_id"`
	TeacherName string            `json:"teacher_name"`
	Subjects    []SubjectResponse `json:"subject"`
}

type AssignSubjectToClassRequest struct {
	SubjectID string `json:"subject_id"`
	TeacherID string `json:"teacher_id"`
	ClassID   string `json:"class_id"`
}

type ClassSubjectResponse struct {
	ClassID     string `json:"class_id"`
	ClassName   string `json:"class_name"`
	SubjectID   string `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	TeacherName string `json:"teacher_name"`
}

type SubjectMattterRequest struct {
	Title       string                        `json:"title"`
	Description string                        `json:"description"`
	Content     []SubjectMatterContentRequest `json:"content"`
}

type SubjectMatterContentRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type SubjectMattterResponse struct {
	ID          string                 `json:"id"`
	Subject     string                 `json:"subject"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
	Content     []SubjectMatterContent `json:"content"`
}

type SubjectMatterContent struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}
