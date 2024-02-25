package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	ID         string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	UserID     string `json:"user_id"`
	User       User   `json:"user" gorm:"foreignKey:UserID"`
	Name       string `json:"name"`
	NISN       string `json:"nisn"`
	Address    string `json:"address"`
	Birthplace string `json:"birthplace"`
	Birthdate  string `json:"birthdate"`
}

func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return nil
}

type Grade struct {
	ID         string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	StudentID  string  `json:"student_id"`
	Student    Student `json:"student" gorm:"foreignKey:StudentID;references:ID"`
	SubjectID  string  `json:"subject_id"`
	Subject    Subject `json:"subject" gorm:"foreignKey:SubjectID;references:ID"`
	TeacherID  string  `json:"teacher_id"`
	Teacher    Teacher `json:"teacher" gorm:"foreignKey:TeacherID;references:ID"`
	Semester   string  `json:"semester"`
	GradeValue float32 `json:"grade_value"`
}

func (g *Grade) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = uuid.NewString()
	return nil
}
