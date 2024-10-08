package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

//go:generate mockgen -source=token_repository.go -destination=mock_token_repository.go -package=mocks
type TokenRepository interface {
	CreateOrUpdateToken(token *entities.Token) (string, error)
	GetTokenByUserID(userID string) (*entities.Token, error)
	FindUserByToken(token string) (string, error)
	GetUserIDByToken(token string) (string, error)
	FindRoleTypeBasedOnToken(token string) (string, error)
	GetTeacherIDByUserID(userID string) (string, error)
	GetStudentIDByUserID(userID string) (string, error)
	DeleteToken(userID string) error
	GetParentIDByUserID(userID string) (string, error)
	GetTeacherByUserID(userID string) (*entities.Teacher, error)
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *tokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) CreateOrUpdateToken(token *entities.Token) (string, error) {
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

func (r *tokenRepository) GetTokenByUserID(userID string) (*entities.Token, error) {
	var token entities.Token
	if err := r.db.Where("user_id = ?", userID).First(&token).Error; err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *tokenRepository) FindUserByToken(token string) (string, error) {
	var user entities.Token
	err := r.db.Where("token = ?", token).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.UserID, nil
}

func (r *tokenRepository) GetUserIDByToken(token string) (string, error) {
	var user entities.Token
	err := r.db.Where("token = ?", token).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.UserID, nil
}

func (r *tokenRepository) FindRoleTypeBasedOnToken(token string) (string, error) {
	var user entities.Token
	err := r.db.Where("token = ?", token).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.RoleType, nil
}

func (r *tokenRepository) GetTeacherIDByUserID(userID string) (string, error) {
	var teacher entities.Teacher
	err := r.db.Where("user_id = ?", userID).First(&teacher).Error
	if err != nil {
		return "", err
	}
	return teacher.ID, nil
}

func (r *tokenRepository) GetStudentIDByUserID(userID string) (string, error) {
	var student entities.Student
	err := r.db.Where("user_id =?", userID).First(&student).Error
	if err != nil {
		return "", err
	}
	return student.ID, nil
}

func (r *tokenRepository) DeleteToken(userID string) error {
	return r.db.Delete(&entities.Token{}, "user_id = ?", userID).Error
}

func (r *tokenRepository) GetParentIDByUserID(userID string) (string, error) {
	var parent entities.Parent
	err := r.db.Where("user_id = ?", userID).First(&parent).Error
	if err != nil {
		return "", err
	}
	return parent.ID, nil
}

func (r *tokenRepository) GetTeacherByUserID(userID string) (*entities.Teacher, error) {
	var teacher entities.Teacher
	err := r.db.Where("user_id = ?", userID).First(&teacher).Error
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}
