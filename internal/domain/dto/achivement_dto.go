package dto

type AchivementRequest struct {
	Title            string `json:"title"`
	TypeOfAchivement string `json:"type_of_achivement"`
	Participation    string `json:"participation"` // winner, participant
	Level            string `json:"level"`         // school, regional, national, international
	Evidence         string `json:"evidence"`
}

type AchivementResponse struct {
	ID               string `json:"id"`
	StudentID        string `json:"student_id"`
	StudentName      string `json:"student_name"`
	Title            string `json:"title"`
	TypeOfAchivement string `json:"type_of_achivement"`
	Participation    string `json:"participation"`
	Level            string `json:"level"`
	Evidence         string `json:"evidence"`
	Status           string `json:"status"`
}

type UpdateAchivementRequest struct {
	Status string `json:"status"`
}
