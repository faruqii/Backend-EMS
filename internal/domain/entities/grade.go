package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Grade struct {
	ID              string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	StudentID       string  `json:"student_id"`
	Student         Student `json:"student" gorm:"foreignKey:StudentID"`
	SubjectID       string  `json:"subject_id"`
	Subject         Subject `json:"subject" gorm:"foreignKey:SubjectID"`
	TeacherID       string  `json:"teacher_id"`
	Teacher         Teacher `json:"teacher" gorm:"foreignKey:TeacherID"`
	Semester        int     `json:"semester"`
	AcademicYear    string  `json:"academic_year"`
	FormativeScores float32 `json:"formative_scores"` // Formative scores
	SummativeScores float32 `json:"summative_scores"` // Summative scores
	ProjectScores   float32 `json:"project_scores"`   // Project scores
	FinalGrade      float32 `json:"final_grade"`      // Final grade
}

func (g *Grade) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = uuid.NewString()
	return nil
}
