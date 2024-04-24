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
	AttendaceAt     time.Time `json:"attendace_at"`
	AttendaceStatus string    `json:"attendace_status"` //izin, sakit,alfa,hadir,terlambat
}

func (a *Atendance) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.NewString()
	return nil
}

type Permits struct {
	ID           string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	StudentID    string    `json:"student_id"`
	Student      Student   `json:"student" gorm:"foreignKey:StudentID"`
	PermitAt     time.Time `json:"permit_at"`
	PermitType   string    `json:"permit_type"`   //izin, sakit
	PermitStatus string    `json:"permit_status"` //diterima, ditolak
	PermitReason string    `json:"permit_reason"`
	Evidence     string    `json:"evidence"` // file type
}

func (p *Permits) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.NewString()
	return nil
}
