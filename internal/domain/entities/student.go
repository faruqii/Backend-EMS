package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	ID          string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	UserID      string  `json:"user_id"`
	User        User    `json:"user" gorm:"foreignKey:UserID"`
	Name        string  `json:"name"`
	NISN        string  `json:"nisn"`
	Gender      string  `json:"gender"` // Laki-laki, Perempuan
	Address     string  `json:"address"`
	Birthplace  string  `json:"birthplace"`
	Birthdate   string  `json:"birthdate"`
	Province    string  `json:"province"`
	City        string  `json:"city"`
	BloodType   string  `json:"blood_type"` // A, B, AB, O
	Religion    string  `json:"religion"`   // Islam, Kristen, Katolik, Hindu, Budha, Konghucu
	Phone       string  `json:"phone"`
	ParentPhone string  `json:"parent_phone"`
	Email       string  `json:"email"`
	ClassID     *string `json:"class_id"`
	Class       Class   `json:"class" gorm:"foreignKey:ClassID"`
}

func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return nil
}
