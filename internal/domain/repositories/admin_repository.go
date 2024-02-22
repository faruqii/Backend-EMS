package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByUsername(username string) (*entities.Admin, error)
	FindById(id string) (*entities.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) FindByUsername(username string) (*entities.Admin, error) {
	var admin entities.Admin
	err := r.db.Where("username = ?", username).First(&admin).Error
	return &admin, err
}

func (r *adminRepository) FindById(id string) (*entities.Admin, error) {
	var admin entities.Admin
	err := r.db.Where("id = ?", id).First(&admin).Error
	return &admin, err
}

// Path: internal/domain/repositories/student_repository.go
