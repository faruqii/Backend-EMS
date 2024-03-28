package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Grade struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	StudentID  uuid.UUID `json:"student_id"`
	Student    Student   `json:"student" gorm:"foreignKey:StudentID"`
	SubjectID  uuid.UUID `json:"subject_id"`
	Subject    Subject   `json:"subject" gorm:"foreignKey:SubjectID"`
	TeacherID  uuid.UUID `json:"teacher_id"`
	Teacher    Teacher   `json:"teacher" gorm:"foreignKey:TeacherID"`
	Semester   int       `json:"semester"`
	FinalGrade float32   `json:"final_grade"` // Final grade
}

func (g *Grade) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = uuid.New()
	return nil
}
