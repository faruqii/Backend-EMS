package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type AdminRepository interface {
	Insert(admin *entities.Admin) (*entities.Admin, error)
	Update(admin *entities.Admin) (*entities.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db:db}
}

/**
Super Admin Can Insert Admin
**/
func (r *adminRepository) Insert(admin *entities.Admin) (*entities.Admin, error) {
	err := r.db.Create(admin).Error

	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *adminRepository) Update(admin *entities.Admin) (*entities.Admin, error) {
	err := r.db.Save(admin).Error

	if err != nil {
		return nil, err
	}

	return admin, nil
}
