package dto

type LiterationRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Documents   string `json:"documents"`
}

type LiterationResponse struct {
	ID          string `json:"id"`
	StudentID   string `json:"student_id"`
	Student     string `json:"student"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Documents   string `json:"documents"`
	Feedback    string `json:"feedback"`
}

type UpdateLiterationRequest struct {
	Feedback string `json:"feedback"`
}
