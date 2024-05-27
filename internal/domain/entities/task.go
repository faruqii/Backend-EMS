package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID          string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	ClassID     string    `json:"class_id"`
	Class       Class     `json:"class" gorm:"foreignKey:ClassID"`
	SubjectID   string    `json:"subject_id"`
	Subject     Subject   `json:"subject" gorm:"foreignKey:SubjectID"`
	TeacherID   string    `json:"teacher_id"`
	Teacher     Teacher   `json:"teacher" gorm:"foreignKey:TeacherID"`
	Title       string    `json:"title"`
	TypeOfTask  string    `json:"type_of_task"`
	Description string    `json:"description"`
	Deadline    string    `json:"deadline"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.NewString()
	return nil
}
