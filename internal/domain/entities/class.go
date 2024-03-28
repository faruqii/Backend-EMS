package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	ID                string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name              string  `json:"name"`
	HomeRoomTeacherID *string `json:"homeRoomTeacherID"`
	HomeRoomTeacher   Teacher `json:"homeRoomTeacher" gorm:"foreignKey:HomeRoomTeacherID;references:ID"`
}

func (c *Class) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.NewString()
	return nil
}
