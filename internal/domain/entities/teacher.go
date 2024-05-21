package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Teacher struct {
	ID          string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	UserID      string    `json:"user_id"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	IsHomeroom  bool      `json:"isHomeroom"`
	IsCouncelor bool      `json:"isCouncelor"`
	Subjects    []Subject `json:"subject" gorm:"many2many:teacher_subjects"`
}

func (t *Teacher) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.NewString()
	return nil
}
