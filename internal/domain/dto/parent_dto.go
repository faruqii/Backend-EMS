package dto

type ParentRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Occupation  string `json:"occupation"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type ParentResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Occupation  string `json:"occupation"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	StudentName string `json:"student_name"`
}

type ParentStudentRequest struct {
	ParentID  string `json:"parent_id"`
	StudentID string `json:"student_id"`
}

type ParentStudentResponse struct {
	ParentID  string `json:"parent_id"`
	StudentID string `json:"student_id"`
}

type ParentProfileResponse struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Occupation  string `json:"occupation"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
