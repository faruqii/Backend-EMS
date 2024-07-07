package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subject struct {
	ID            string         `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Semester      string         `json:"semester"`
	Teachers      []Teacher      `json:"teachers" gorm:"many2many:teacher_subjects"`
	ClassSubjects []ClassSubject `json:"class_subjects" gorm:"foreignKey:SubjectID"`
}

func (s *Subject) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return nil
}

type TeacherSubject struct {
	TeacherID string  `json:"teacher_id"`
	SubjectID string  `json:"subject_id"`
	Teacher   Teacher `json:"teacher" gorm:"foreignKey:TeacherID"`
	Subject   Subject `json:"subject" gorm:"foreignKey:SubjectID"`
}

type ClassSubject struct {
	ID        string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	ClassID   string  `json:"class_id"`
	Class     Class   `json:"class" gorm:"foreignKey:ClassID"`
	SubjectID string  `json:"subject_id"`
	Subject   Subject `json:"subject" gorm:"foreignKey:SubjectID"`
	TeacherID string  `json:"teacher_id"`
	Teacher   Teacher `json:"teacher" gorm:"foreignKey:TeacherID"`
}

func (cs *ClassSubject) BeforeCreate(tx *gorm.DB) (err error) {
	cs.ID = uuid.NewString()
	return nil
}

type SubjectMattter struct {
    ID          string                 `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
    SubjectID   string                 `json:"subject_id"`
    Subject     Subject                `json:"subject" gorm:"foreignKey:SubjectID"`
    Title       string                 `json:"title"`
    Description string                 `json:"description"`
    Content     []SubjectMatterContent `json:"content" gorm:"foreignKey:SubjectMatterID"`
}

type SubjectMatterContent struct {
    ID              string          `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
    SubjectMatterID string          `json:"subject_matter_id" gorm:"type:uuid"`
    SubjectMatter   *SubjectMattter `json:"subject_matter,omitempty" gorm:"foreignKey:SubjectMatterID;references:ID"`
    Title           string          `json:"title"`
    Description     string          `json:"description"`
    Link            string          `json:"link"`
}


func (sm *SubjectMattter) BeforeCreate(tx *gorm.DB) (err error) {
	sm.ID = uuid.NewString()
	return nil
}

func (smc *SubjectMatterContent) BeforeCreate(tx *gorm.DB) (err error) {
	smc.ID = uuid.NewString()
	return nil
}
