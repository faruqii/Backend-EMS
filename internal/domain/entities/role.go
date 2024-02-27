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

// table users
// 1, "name", "pass"
// table roles
// 1, "superadmin"
// 2, "admin"
// 3, "teacher"
// 4, "student"
// 5, "parent"
// table user_roles (map user_id to role_id)
// 1, 1
// 2, 2
// 3, 3
// table tokens
// id, user_id, token, role_type, expired_at

// authentication design
// 1. user logs in with username and password
// 2. server checks if user exists
// 3. server checks if password matches
// 4. server returns a token

// belong to authentication middleware
// 5. user uses token to access protected routes
// 6. server checks if token is valid

// belongs to authorization middleware
// 7. server checks user's role
// 8. server checks if user has access to the route
// 9. server returns data to user
