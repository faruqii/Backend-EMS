package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	ID                string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name              string  `json:"name"`
	HomeRoomTeacherID *string `json:"homeRoomTeacherID"`
	HomeRoomTeacher   *Teacher
	Students          []Student `json:"students" gorm:"many2many:student_classes"`
}

func (c *Class) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.NewString()
	return nil
}

type Schedule struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	ClassID   string    `json:"class_id"`
	Class     Class     `gorm:"foreignKey:ClassID" json:"class"`
	SubjectID string    `json:"subject_id"`
	Subject   Subject   `gorm:"foreignKey:SubjectID" json:"subject"`
	TeacherID string    `json:"teacher_id"`
	Teacher   Teacher   `gorm:"foreignKey:TeacherID" json:"teacher"`
	DayOfWeek string    `json:"day_of_week"` // Consider using an enum or constant for days
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (s *Schedule) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return nil
}
