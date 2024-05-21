package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Counseling struct {
	ID        string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	StudentID string  `json:"student_id"`
	Student   Student `json:"student" gorm:"foreignKey:StudentID"`
	Purpose   string  `json:"purpose"` // tujuan konseling : akademik, non-akademik
	Type      string  `json:"type"`    // tipe konseling : baru, lanjutan
	Date      string  `json:"date"`    // format: dd-mm-yyyy
	Status    string  `json:"status"`  // status konseling : Wait Approval, Approved, Declined
}

func (c *Counseling) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.NewString()
	return nil
}
