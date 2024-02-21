package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type TeacherRepository interface {
	FindByID(id string) (*entities.Teacher, error)
	FindByEmail(email string) (*entities.Teacher, error)
	Create(teacher *entities.Teacher) error
	Update(teacher *entities.Teacher) error
	Delete(id string) error
	GetAll() ([]entities.Teacher, error)
}

// teacherRepository is a concrete implementation of TeacherRepository.
type teacherRepository struct {
	db *gorm.DB
}

// NewTeacherRepository creates a new instance of teacherRepository.
func NewTeacherRepository(db *gorm.DB) TeacherRepository {
	return &teacherRepository{db: db}
}

// FindByID finds a teacher by ID.
func (r *teacherRepository) FindByID(id string) (*entities.Teacher, error) {
	var teacher entities.Teacher
	if err := r.db.First(&teacher, id).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

// FindByEmail finds a teacher by email.
func (r *teacherRepository) FindByEmail(email string) (*entities.Teacher, error) {
	var teacher entities.Teacher
	if err := r.db.Where("email = ?", email).First(&teacher).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

// Create creates a new teacher.
func (r *teacherRepository) Create(teacher *entities.Teacher) error {
	if err := r.db.Create(teacher).Error; err != nil {
		return err
	}
	return nil
}

// Update updates an existing teacher.
func (r *teacherRepository) Update(teacher *entities.Teacher) error {
	if err := r.db.Save(teacher).Error; err != nil {
		return err
	}
	return nil
}

// Delete deletes a teacher by ID.
func (r *teacherRepository) Delete(id string) error {
	if err := r.db.Delete(&entities.Teacher{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetAll retrieves all teachers.
func (r *teacherRepository) GetAll() ([]entities.Teacher, error) {
	var teachers []entities.Teacher
	if err := r.db.Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}
