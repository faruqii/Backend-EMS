package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)
//go:generate mockgen -source=role_repository.go -destination=mock_role_repository.go -package=mocks
type RoleRepository interface {
	GetRoleNameFromID(id string) (string, error)
	AssignUserRole(userID, roleName string) error
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

func (r *roleRepository) AssignUserRole(userID, roleName string) error {
	var role entities.Role
	if err := r.db.Where("name = ?", roleName).First(&role).Error; err != nil {
		return err
	}
	userRole := entities.UserRole{
		UserID: userID,
		RoleID: role.ID,
	}
	if err := r.db.Create(&userRole).Error; err != nil {
		return err
	}
	return nil
}
