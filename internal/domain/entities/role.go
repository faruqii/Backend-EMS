package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	SuperAdminRole Role = "superadmin"
	AdminRole      Role = "admin"
	StudentRole    Role = "student"
	TeacherRole    Role = "teacher"
	ParentRole     Role = "parent"
)

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

type Student struct {
	User
	Name       string `json:"name"`
	NISN       string `json:"nisn"`
	Address    string `json:"address"`
	Birthplace string `json:"birthplace"`
	Birthdate  string `json:"birthdate"`
	ParentID   string `json:"parent_id"`
	Parent     Parent `gorm:"foreignKey:ParentID"`
}

type Teacher struct {
	User
	Name       string `json:"name"`
	NIP        string `json:"nip"`
	Address    string `json:"address"`
	Birthplace string `json:"birthplace"`
	Birthdate  string `json:"birthdate"`
}

type Parent struct {
	User
	Name       string `json:"name"`
	Address    string `json:"address"`
	Birthplace string `json:"birthplace"`
	Birthdate  string `json:"birthdate"`
}
