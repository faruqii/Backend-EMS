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
