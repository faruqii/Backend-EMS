package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Announcement struct {
	ID          string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Title       string    `json:"title"`
	Information string    `json:"information"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (a *Announcement) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.NewString()
	return nil
}
