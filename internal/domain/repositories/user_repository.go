package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entities.User) error
	Update(user *entities.User) error
	Delete(id string) error
	FindByUsername(username string) (*entities.User, error)
	FindById(id string) (*entities.User, error)
	ChangePassword(userID string, newPassword string) error
	IsPasswordMatch(userID string, password string) bool
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id string) error {
	return r.db.Delete(&entities.User{}, id).Error
}

func (r *userRepository) FindByUsername(username string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *userRepository) FindById(id string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *userRepository) ChangePassword(userID string, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return r.db.Model(&entities.User{}).Where("id = ?", userID).Update("password", string(hashedPassword)).Error
}

func (r *userRepository) IsPasswordMatch(userID string, password string) bool {
	var user entities.User
	r.db.Where("id = ?", userID).First(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// Path: internal/domain/repositories/student_repository.go
