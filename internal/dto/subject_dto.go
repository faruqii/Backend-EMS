package dto

type SubjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Semester    string `json:"semester"`
}

type SubjectResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Semester    string `json:"semester"`
	// Omitting Teachers field as it's not included in the response
}
