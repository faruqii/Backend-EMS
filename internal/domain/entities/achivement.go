package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentAchivement struct {
	ID               string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	StudentID        string  `json:"student_id"`
	Student          Student `json:"student" gorm:"foreignKey:StudentID"`
	Title            string  `json:"title"`
	TypeOfAchivement string  `json:"type_of_achivement"`
	Participation    string  `json:"participation"` // winner, participant
	Level            string  `json:"level"`         // school, regional, national, international
	Evidence         string  `json:"evidence"`
	Status           string  `json:"status"` // pending, approved, rejected
}

func (sa *StudentAchivement) BeforeCreate(tx *gorm.DB) (err error) {
	sa.ID = uuid.NewString()
	return nil
}
