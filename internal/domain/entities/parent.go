package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Parent struct {
	ID          string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	UserID      string `json:"user_id"`
	User        User   `json:"user" gorm:"foreignKey:UserID"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Occupation  string `json:"occupation"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

func (p *Parent) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.NewString()
	return nil
}

type ParentStudent struct {
	ID        string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	ParentID  string  `json:"parent_id"`
	StudentID string  `json:"student_id"`
	Parent    Parent  `json:"parent" gorm:"foreignKey:ParentID"`
	Student   Student `json:"student" gorm:"foreignKey:StudentID"`
}

func (pc *ParentStudent) BeforeCreate(tx *gorm.DB) (err error) {
	pc.ID = uuid.NewString()
	return nil
}
