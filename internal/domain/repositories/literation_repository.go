package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type LiterationRepository interface {
	InsertLiteration(literation *entities.Literation) (*entities.Literation, error)
	GetLiterationByID(literationID string) (*entities.Literation, error)
	GetLiterationByStudentID(studentID string) ([]entities.Literation, error)
	GetAllLiterations() ([]entities.Literation, error)
	Update(literationID string, feedback string, point int, status string) (*entities.Literation, error)
	FilterByStudentClass(classID string) ([]entities.Literation, error)
}

type literationRepository struct {
	db *gorm.DB
}

func NewLiterationRepository(db *gorm.DB) LiterationRepository {
	return &literationRepository{
		db: db,
	}
}

func (r *literationRepository) InsertLiteration(literation *entities.Literation) (*entities.Literation, error) {
	if err := r.db.Create(literation).Error; err != nil {
		return nil, err
	}
	return literation, nil
}

func (r *literationRepository) GetLiterationByID(literationID string) (*entities.Literation, error) {
	var literation entities.Literation
	// preloading the student data
	if err := r.db.Preload("Student").Preload("Student.Class").Where("id = ?", literationID).First(&literation).Error; err != nil {
		return nil, err
	}
	return &literation, nil
}

func (r *literationRepository) GetLiterationByStudentID(studentID string) ([]entities.Literation, error) {
	var literations []entities.Literation
	// preloading the student data
	if err := r.db.Preload("Student").Preload("Student.Class").Where("student_id = ?", studentID).Find(&literations).Error; err != nil {
		return nil, err
	}
	return literations, nil
}

func (r *literationRepository) GetAllLiterations() ([]entities.Literation, error) {
	var literations []entities.Literation
	// preloading the student data
	if err := r.db.Preload("Student").Preload("Student.Class").Find(&literations).Error; err != nil {
		return nil, err
	}
	return literations, nil
}

func (r *literationRepository) Update(literationID string, feedback string, point int, status string) (*entities.Literation, error) {
	var literation entities.Literation
	if err := r.db.Where("id = ?", literationID).First(&literation).Error; err != nil {
		return nil, err
	}
	literation.Feedback = feedback
	literation.Points = point
	literation.Status = status
	if err := r.db.Save(&literation).Error; err != nil {
		return nil, err
	}
	return &literation, nil
}

func (r *literationRepository) FilterByStudentClass(classID string) ([]entities.Literation, error) {
	var literations []entities.Literation
	// preloading the student data
	if err := r.db.Preload("Student").Joins("JOIN students ON literations.student_id = students.id").Where("students.class_id = ?", classID).Find(&literations).Error; err != nil {
		return nil, err
	}
	return literations, nil
}
