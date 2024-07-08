package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Agenda struct {
	ID           string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Title        string    `json:"title"`
	Date         time.Time `json:"date"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	TypeOfAgenda string    `json:"type_of_agenda"`
	Location     string    `json:"location"`
	Description  string    `json:"description"`
}

func (a *Agenda) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.NewString()
	return nil
}
