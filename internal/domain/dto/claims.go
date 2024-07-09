package dto

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserID            string `json:"user_id"`
	Role              string `json:"role"`
	IsHomeroomTeacher bool   `json:"is_homeroom_teacher,omitempty"`
	jwt.RegisteredClaims
}

