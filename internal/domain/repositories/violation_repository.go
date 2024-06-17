package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type ViolationRepository interface {
	Create(violation *entities.Violation) error
	GetAll() ([]entities.Violation, error)
	GetByID(id string) (*entities.Violation, error)
	GetByStudentID(studentID string) ([]entities.Violation, error)
	Update(violation *entities.Violation) error
	Delete(id string) error
}

type violationRepository struct {
	db *gorm.DB
}

func NewViolationRepository(db *gorm.DB) ViolationRepository {
	return &violationRepository{db: db}
}

func (r *violationRepository) Create(violation *entities.Violation) error {
	return r.db.Create(violation).Error
}

func (r *violationRepository) GetAll() ([]entities.Violation, error) {
	var violations []entities.Violation
	err := r.db.Preload("Student").Find(&violations).Error
	return violations, err
}

func (r *violationRepository) GetByID(id string) (*entities.Violation, error) {
	var violation entities.Violation
	err := r.db.Preload("Student").Where("id = ?", id).First(&violation).Error
	return &violation, err
}

func (r *violationRepository) GetByStudentID(studentID string) ([]entities.Violation, error) {
	var violations []entities.Violation
	err := r.db.Preload("Student").Where("student_id = ?", studentID).Find(&violations).Error
	return violations, err
}

func (r *violationRepository) Update(violation *entities.Violation) error {
	return r.db.Save(violation).Error
}

func (r *violationRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&entities.Violation{}).Error
}
