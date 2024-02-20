package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByUsername(username string) (*entities.Admin, error)
	FindById(id string) (*entities.Admin, error)
	CreateOrUpdateToken(token *entities.Token) (string, error)
	GetTokenByUserID(userID string) (*entities.Token, error)
	FindAdminByToken(token string) (string, error)
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

func (r *adminRepository) FindAdminByToken(token string) (string, error) {
	var admin entities.Admin
	err := r.db.Where("token = ?", token).First(&admin).Error
	return admin.Username, err
}

func (r *adminRepository) CreateOrUpdateToken(token *entities.Token) (string, error) {
	var existingToken entities.Token
	err := r.db.Where("user_id = ?", token.UserID).First(&existingToken).Error

	if err != nil {
		if err := r.db.Create(token).Error; err != nil {
			return "", err
		}
	} else {
		token.ID = existingToken.ID
		if err := r.db.Save(token).Error; err != nil {
			return "", err
		}
	}

	return token.Token, nil
}

func (r *adminRepository) GetTokenByUserID(userID string) (*entities.Token, error) {
	var token entities.Token
	if err := r.db.Where("user_id = ?", userID).First(&token).Error; err != nil {
		return nil, err
	}

	return &token, nil
}

// Path: internal/domain/repositories/student_repository.go
