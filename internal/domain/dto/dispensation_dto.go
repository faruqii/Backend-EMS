package dto

type DispensationRequest struct {
	Reason   string `json:"reason"`
	StartAt  string `json:"start_at"`
	EndAt    string `json:"end_at"`
	Document string `json:"document"`
}

type DispensationResponse struct {
	ID        string `json:"id"`
	StudentID string `json:"student_id"`
	Student   string `json:"student"`
	Reason    string `json:"reason"`
	StartAt   string `json:"start_at"`
	EndAt     string `json:"end_at"`
	Document  string `json:"document"`
	Status    string `json:"status"`
}
