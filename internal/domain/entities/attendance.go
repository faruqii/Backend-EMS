package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Atendance struct {
	ID              string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	StudentID       string    `json:"student_id"`
	Student         Student   `json:"student" gorm:"foreignKey:StudentID"`
	SubjectID       string    `json:"subject_id"`
	Subject         Subject   `json:"subject" gorm:"foreignKey:SubjectID"`
	AttendaceStatus string    `json:"attendace_status"` //izin, sakit,alfa,hadir,terlambat
	AttendaceAt     time.Time `json:"attendace_at"`
}

func (a *Atendance) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.NewString()
	return nil
}
