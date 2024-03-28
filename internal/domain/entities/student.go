package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	ID         string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	UserID     string  `json:"user_id"`
	User       User    `json:"user" gorm:"foreignKey:UserID"`
	Name       string  `json:"name"`
	NISN       string  `json:"nisn"`
	Address    string  `json:"address"`
	Birthplace string  `json:"birthplace"`
	Birthdate  string  `json:"birthdate"`
	ClassID    *string `json:"class_id"`
	Class      Class   `json:"class" gorm:"foreignKey:ClassID"`
}

func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return nil
}
