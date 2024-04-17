package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return nil
}

type UserRole struct {
	UserID string `json:"user_id" gorm:"primaryKey"`
	RoleID string `json:"role_id" gorm:"primaryKey"`
	User   User
	Role   Role
}

