package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID          string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	SubjectID   string  `json:"subject_id"`
	Subject     Subject `json:"subject" gorm:"foreignKey:SubjectID"`
	TeacherID   string  `json:"teacher_id"`
	Teacher     Teacher `json:"teacher" gorm:"foreignKey:TeacherID"`
	Title       string  `json:"title"`
	TypeOfTask  string  `json:"type_of_task"` // tugas harian, ulangan harian, ulangan tengah semester, ulangan akhir semester, quiz, dll
	Description string  `json:"description"`
	Deadline    string  `json:"deadline"` // format: dd-mm-yyyy
	Link        string  `json:"link"`     // link to google drive or something else
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.NewString()
	return nil
}
