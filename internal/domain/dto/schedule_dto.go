package dto

import "time"

type CreateScheduleRequest struct {
	ClassID   string       `json:"class_id"`
	SubjectID string       `json:"subject_id"`
	TeacherID string       `json:"teacher_id"`
	DayOfWeek time.Weekday `json:"day_of_week"`
	StartTime string       `json:"start_time"`
	EndTime   string       `json:"end_time"`
}

type ScheduleResponse struct {
	ID        string `json:"id"`
	Class     string `json:"class"`
	Subject   string `json:"subject"`
	Teacher   string `json:"teacher"`
	DayOfWeek string `json:"day_of_week"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
