package dto

type LiterationRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Documents   string `json:"documents"`
}

type LiterationResponse struct {
	ID             string `json:"id"`
	StudentID      string `json:"student_id"`
	Student        string `json:"student"`
	StudentClassID string `json:"student_class_id"`
	StudentClass   string `json:"student_class"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Documents      string `json:"documents"`
	Feedback       string `json:"feedback"`
	Point          int    `json:"point"`
	Status         string `json:"status"`
}

type UpdateLiterationRequest struct {
	Feedback string `json:"feedback"`
	Point    int    `json:"point"`
	Status   string `json:"status"`
}
