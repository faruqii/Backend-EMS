package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Violation struct {
	ID              string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	StudentID       string    `json:"student_id"`
	Student         Student   `json:"student" gorm:"foreignKey:StudentID"`
	SKNumber        string    `json:"sk"`
	StartPunishment time.Time `json:"start_punishment"`
	EndPunishment   time.Time `json:"end_punishment"`
	Documents       string    `json:"documents"`
	Reason          string    `json:"reason"`
}

func (v *Violation) BeforeCreate(tx *gorm.DB) (err error) {
	v.ID = uuid.NewString()
	return nil
}
