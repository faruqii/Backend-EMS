package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Literation struct {
	ID          string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	StudentID   string  `json:"student_id"`
	Student     Student `json:"student" gorm:"foreignKey:StudentID"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Documents   string  `json:"documents"`
	Feedback    string  `json:"feedback"`
	Points      int     `json:"points"`
	Status      string  `json:"status"`
}

func (l *Literation) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.NewString()
	return nil
}
