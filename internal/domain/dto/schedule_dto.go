package dto

import "time"

type CreateScheduleRequest struct {
	ClassID   string        `json:"class_id"`
	SubjectID string        `json:"subject_id"`
	TeacherID string        `json:"teacher_id"`
	DayOfWeek time.Weekday  `json:"day_of_week"`
	Duration  time.Duration `json:"duration"`
}

type ScheduleResponse struct {
	ID        string        `json:"id"`
	Class     string        `json:"class"`
	Subject   string        `json:"subject"`
	Teacher   string        `json:"teacher"`
	DayOfWeek time.Weekday  `json:"day_of_week"`
	Duration  time.Duration `json:"duration"`
}
