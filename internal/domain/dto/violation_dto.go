package dto

type ViolationRequest struct {
	StudentID       string `json:"student_id"`
	SKNumber        string `json:"sk"`
	StartPunishment string `json:"start_punishment"`
	EndPunishment   string `json:"end_punishment"`
	Documents       string `json:"documents"`
	Reason          string `json:"reason"`
}

type ViolationResponse struct {
	ID              string `json:"id"`
	StudentID       string `json:"student_id"`
	Student         string `json:"student"`
	SKNumber        string `json:"sk"`
	StartPunishment string `json:"start_punishment"`
	EndPunishment   string `json:"end_punishment"`
	Documents       string `json:"documents"`
	Reason          string `json:"reason"`
}
