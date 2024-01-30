package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subject struct {
	ID          string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Semester    string  `json:"semester"`
	TeacherID   string  `json:"teacher_id"`
	Teacher     Teacher `gorm:"foreignKey:TeacherID"`
}

func (s *Subject) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return nil
}
