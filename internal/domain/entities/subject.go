package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subject struct {
	ID          string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Semester    string    `json:"semester"`
	Teachers    []Teacher `gorm:"many2many:teacher_subjects;"`
}

func (s *Subject) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return nil
}

type TeacherSubject struct {
	TeacherID string `gorm:"primaryKey"`
	SubjectID string `gorm:"primaryKey"`
	Teacher   Teacher
	Subject   Subject
}
