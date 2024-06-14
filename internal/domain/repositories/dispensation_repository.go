package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type DispensationRepository interface {
	InsertDispensation(dispensation *entities.Dispensation) (*entities.Dispensation, error)
	GetDispensationByID(id string) (*entities.Dispensation, error)
	GetAllDispensations() ([]entities.Dispensation, error)
	GetDispensationsByStudentID(studentID string) ([]entities.Dispensation, error)
	UpdateDispensationStatus(dispensationID string, status string) (*entities.Dispensation, error)
}

type dispensationRepository struct {
	db *gorm.DB
}

func NewDispensationRepository(db *gorm.DB) DispensationRepository {
	return &dispensationRepository{db: db}
}

func (r *dispensationRepository) InsertDispensation(dispensation *entities.Dispensation) (*entities.Dispensation, error) {
	if err := r.db.Create(&dispensation).Error; err != nil {
		return nil, err
	}
	return dispensation, nil
}

func (r *dispensationRepository) GetDispensationByID(id string) (*entities.Dispensation, error) {
	// preload student
	var dispensation entities.Dispensation
	if err := r.db.Preload("Student").Where("id = ?", id).First(&dispensation).Error; err != nil {
		return nil, err
	}

	return &dispensation, nil
}

func (r *dispensationRepository) GetAllDispensations() ([]entities.Dispensation, error) {
	// preload student
	var dispensations []entities.Dispensation
	if err := r.db.Preload("Student").Find(&dispensations).Error; err != nil {
		return nil, err
	}

	return dispensations, nil
}

func (r *dispensationRepository) GetDispensationsByStudentID(studentID string) ([]entities.Dispensation, error) {
	// preload student
	var dispensations []entities.Dispensation
	if err := r.db.Preload("Student").Where("student_id = ?", studentID).Find(&dispensations).Error; err != nil {
		return nil, err
	}

	return dispensations, nil
}

func (r *dispensationRepository) UpdateDispensationStatus(dispensationID string, status string) (*entities.Dispensation, error) {
	var dispensation entities.Dispensation
	if err := r.db.Model(&dispensation).Where("id = ?", dispensationID).Update("status", status).Error; err != nil {
		return nil, err
	}

	return &dispensation, nil
}
