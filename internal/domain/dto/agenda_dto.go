package dto

type CreateAgendaRequest struct {
	Title        string `json:"title" binding:"required"`
	Date         string `json:"date" binding:"required"`
	StartTime    string `json:"start_time" binding:"required"`
	EndTime      string `json:"end_time" binding:"required"`
	TypeOfAgenda string `json:"type_of_agenda" binding:"required"`
	Location     string `json:"location" binding:"required"`
	Description  string `json:"description"`
}

type UpdateAgendaRequest struct {
	Title        string `json:"title"`
	Date         string `json:"date"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	TypeOfAgenda string `json:"type_of_agenda"`
	Location     string `json:"location"`
	Description  string `json:"description"`
}

type AgendaResponse struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Date         string `json:"date"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	TypeOfAgenda string `json:"type_of_agenda"`
	Location     string `json:"location"`
	Description  string `json:"description"`
}
