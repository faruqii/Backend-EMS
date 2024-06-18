package dto

type CreateAnnouncementRequest struct {
	Title       string `json:"title"`
	Information string `json:"information"`
}

type UpdateAnnouncementRequest struct {
	Title       string `json:"title"`
	Information string `json:"information"`
}

type AnnouncementResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Information string `json:"information"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
