package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID   string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name string `json:"name"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.NewString()
	return nil
}

type User struct {
	ID       string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return nil
}

type SuperAdmin struct {
	User
}

type Admin struct {
	User
}
