package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Dispensation struct {
	ID        string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	StudentID string    `json:"student_id"`
	Student   Student   `json:"student" gorm:"foreignKey:StudentID"`
	Reason    string    `json:"reason"`   // alasan dispensasi
	StartAt   time.Time `json:"start_at"` // format: dd-mm-yyyy
	EndAt     time.Time `json:"end_at"`   // format: dd-mm-yyyy
	Document  string    `json:"document"` // file type
	Status    string    `json:"status"`   // status dispensasi : Wait Approval, Approved, Declined
}

func (d *Dispensation) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.NewString()
	return nil
}
