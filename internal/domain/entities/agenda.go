package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Agenda struct {
	ID           string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Title        string `json:"title"`
	Date         string `json:"date"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	TypeOfAgenda string `json:"type_of_agenda"`
	Location     string `json:"location"`
	Description  string `json:"description"`
}

func (a *Agenda) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.NewString()
	return nil
}
