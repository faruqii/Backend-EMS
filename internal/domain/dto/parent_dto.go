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
}

type ParentStudentRequest struct {
	ParentID  string `json:"parent_id"`
	StudentID string `json:"student_id"`
}

type ParentStudentResponse struct {
	ID        string `json:"id"`
	ParentID  string `json:"parent_id"`
	StudentID string `json:"student_id"`
}
