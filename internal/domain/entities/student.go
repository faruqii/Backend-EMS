package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	User
	Name       string  `json:"name"`
	NISN       string  `json:"nisn"`
	Address    string  `json:"address"`
	Birthplace string  `json:"birthplace"`
	Birthdate  string  `json:"birthdate"`
	ParentID   *string `json:"parent_id"`
	Parent     *Parent `json:"parent"`
}

type Grade struct {
	ID         string   `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	StudentID  *string  `json:"student_id"`
	Student    *Student `gorm:"foreignKey:StudentID" json:"student"`
	SubjectID  *string  `json:"subject_id"`
	Subject    *Subject `gorm:"foreignKey:SubjectID" json:"subject"`
	TeacherID  *string  `json:"teacher_id"`
	Teacher    *Teacher `gorm:"foreignKey:TeacherID" json:"teacher"`
	Semester   string   `json:"semester"`
	GradeValue float32  `json:"grade_value"`
}

func (g *Grade) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = uuid.NewString()
	return nil
}
