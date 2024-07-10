package dto

import "time"

type CreateAttendanceRequest struct {
	StudentID       string `json:"student_id"`
	AttendaceStatus string `json:"attendace_status"`
}

type AttendanceResponse struct {
	ID              string    `json:"id"`
	StudentID       string    `json:"student_id"`
	StudentName     string    `json:"student_name"`
	SubjectID       string    `json:"subject_id"`
	AttendaceStatus string    `json:"attendace_status"`
	AttendaceAt     time.Time `json:"attendace_at"`
}

type StudentAttedanceResponse struct {
	ID              string    `json:"id"`
	StudentID       string    `json:"student_id"`
	StudentName     string    `json:"student_name"`
	SubjectName     string    `json:"subject_name"`
	AttedanceStatus string    `json:"attedance_status"`
	AttedanceAt     time.Time `json:"attedance_at"`
}

type UpdateAttedanceRequest struct {
	AttedanceStatus string `json:"attedance_status"`
}
