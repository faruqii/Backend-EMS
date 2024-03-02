package dto

type CreateClassRequest struct {
	Name string `json:"name"`
}

type ClassResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	HomeRoomTeacher string `json:"homeRoomTeacher"`
}
