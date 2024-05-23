package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentAssignment struct {
	ID         string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	TaskID     string    `json:"task_id"`
	Task       Task      `json:"task" gorm:"foreignKey:TaskID"`
	StudentID  string    `json:"student_id"`
	Student    Student   `json:"student" gorm:"foreignKey:StudentID"`
	Submission string    `json:"submission"`
	Grade      float64   `json:"grade"`
	Feedback   string    `json:"feedback"`
	SubmitAt   time.Time `json:"submit_at"`
}

func (sa *StudentAssignment) BeforeCreate(tx *gorm.DB) (err error) {
	sa.ID = uuid.NewString()
	return nil
}
