package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Quiz struct {
	ID          string     `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	ClassID     string     `json:"class_id"`
	Class       Class      `json:"class" gorm:"foreignKey:ClassID"`
	SubjectID   string     `json:"subject_id"`
	Subject     Subject    `json:"subject" gorm:"foreignKey:SubjectID"`
	TeacherID   string     `json:"teacher_id"`
	Teacher     Teacher    `json:"teacher" gorm:"foreignKey:TeacherID"`
	Title       string     `json:"title"`
	TypeOfQuiz  string     `json:"type_of_quiz"` // Quiz, UTS, UAS
	Description string     `json:"description"`
	Deadline    string     `json:"deadline"`
	Questions   []Question `json:"questions" gorm:"foreignKey:QuizID"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Question struct {
	ID            string         `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	QuizID        string         `json:"quiz_id"`
	Quiz          Quiz           `json:"quiz" gorm:"foreignKey:QuizID"`
	Text          string         `json:"text"`
	Options       pq.StringArray `json:"options" gorm:"type:varchar(255)[]"`
	CorrectAnswer string         `json:"correct_answer"`
}

func (q *Quiz) BeforeCreate(tx *gorm.DB) (err error) {
	q.ID = uuid.NewString()
	return nil
}

func (q *Question) BeforeCreate(tx *gorm.DB) (err error) {
	q.ID = uuid.NewString()
	return nil
}
