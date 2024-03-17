package dto

type TeacherRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type TeacherResponse struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	IsHomeroomTeacher bool   `json:"is_homeroom_teacher"`
}

type UpdateHomeroomStatusRequest struct {
	Status bool `json:"status"`
}

type TeacherSchedule struct {
	SubjectName string `json:"subject_name"`
	ClassName   string `json:"class_name"`
	Day         string `json:"day"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}
