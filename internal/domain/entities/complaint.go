package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Complaint struct {
	ID          string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	UserID      string    `json:"user_id"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	Description string    `json:"description"`
	Documents   string    `json:"documents"` // file type
	ComplaintAt time.Time `json:"complaint_at"`
	Status      string    `json:"status"` 
}

func (c *Complaint) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.NewString()
	return nil
}
