package repositories

import (
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type DispensationRepository interface {
	InsertDispensation(dispensation *entities.Dispensation) (*entities.Dispensation, error)
	GetDispensationByID(id string) (*entities.Dispensation, error)
	GetAllDispensations() ([]entities.Dispensation, error)
	GetDispensationsByStudentID(studentID string) ([]entities.Dispensation, error)
	UpdateDispensation(dispensation *entities.Dispensation) (*entities.Dispensation, error)
	ShowTodayDispensation() ([]entities.Dispensation, error)
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

func (r *dispensationRepository) UpdateDispensation(dispensation *entities.Dispensation) (*entities.Dispensation, error) {
	var oldDispensation entities.Dispensation
	if err := r.db.Where("id = ?", dispensation.ID).First(&oldDispensation).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&oldDispensation).Updates(dispensation).Error; err != nil {
		return nil, err
	}

	return dispensation, nil
}

func (r *dispensationRepository) ShowTodayDispensation() ([]entities.Dispensation, error) {
	var dispensations []entities.Dispensation

	// Get today's date in dd-mm-yyyy format
	today := time.Now().Format("02-01-2006")

	if err := r.db.Preload("Student").Where("start_at <= ? AND end_at >= ?", today, today).Find(&dispensations).Error; err != nil {
		return nil, err
	}

	return dispensations, nil
}
