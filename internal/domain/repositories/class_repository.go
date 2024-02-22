package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type ClassRepository interface {
	Insert(class *entities.Class) error
	Update(class *entities.Class) error
	Delete(id string) error
	FindByID(id string) (*entities.Class, error)
	FindByTeacherID(teacherID string) ([]entities.Class, error)
	GetAll() ([]entities.Class, error)
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{db: db}
}

func (r *classRepository) Insert(class *entities.Class) error {
	if err := r.db.Create(class).Error; err != nil {
		return err
	}
	return nil
}

func (r *classRepository) Update(class *entities.Class) error {
	if err := r.db.Save(class).Error; err != nil {
		return err
	}
	return nil
}

func (r *classRepository) Delete(id string) error {
	if err := r.db.Delete(&entities.Class{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *classRepository) FindByID(id string) (*entities.Class, error) {
	var class entities.Class
	if err := r.db.First(&class, id).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

func (r *classRepository) FindByTeacherID(teacherID string) ([]entities.Class, error) {
	var classes []entities.Class
	if err := r.db.Where("teacher_id = ?", teacherID).Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

func (r *classRepository) GetAll() ([]entities.Class, error) {
	var classes []entities.Class
	if err := r.db.Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}
