package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	ID                string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name              string    `json:"name"`
	HomeRoomTeacherID *string   `json:"homeRoomTeacherID"`
	HomeRoomTeacher   Teacher   `json:"homeRoomTeacher" gorm:"foreignKey:HomeRoomTeacherID;references:ID"`
	Students          []Student `json:"students" gorm:"many2many:student_classes"`
}

func (c *Class) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.NewString()
	return nil
}

type Schedule struct {
	ID        string       `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	ClassID   string       `json:"class_id"`
	Class     Class        `json:"class" gorm:"foreignKey:ClassID;references:ID"`
	SubjectID string       `json:"subject_id"`
	Subject   Subject      `json:"subject" gorm:"foreignKey:SubjectID;references:ID"`
	TeacherID string       `json:"teacher_id"`
	Teacher   Teacher      `json:"teacher" gorm:"foreignKey:TeacherID;references:ID"`
	DayOfWeek time.Weekday `json:"day_of_week"` // Day starts from 0 (Sunday) to 6 (Saturday)
	StartTime time.Time    `json:"start_time"`
	EndTime   time.Time    `json:"end_time"`
}

func (s *Schedule) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return nil
}
