package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Agenda struct {
	ID          string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Date        string `json:"date"`
	StartAt     string `json:"start_at"`
	EndAt       string `json:"end_at"`
}

func (a *Agenda) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.NewString()
	return nil
}
