package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRoleNameFromID(id string) (string, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) GetRoleNameFromID(userID string) (string, error) {
	var userRole entities.UserRole
	if err := r.db.Preload("Role").Where("user_id = ?", userID).First(&userRole).Error; err != nil {
		return "", err
	}
	return userRole.Role.Name, nil
}
