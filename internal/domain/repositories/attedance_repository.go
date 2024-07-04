package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type AttedanceRepository interface {
	CreateAttedance(attedance *entities.Atendance) (*entities.Atendance, error)
	GetAttedanceBySubjectID(subjectID string) ([]entities.Atendance, error)
	GetAttedanceByClassID(classID string) ([]entities.Atendance, error)
	GetAttedanceBySubjectAndClassID(subjectID, classID string) ([]entities.Atendance, error)
	GetMyAttedance(studentID string) ([]entities.Atendance, error)
	UpdateAttedance(attedance *entities.Atendance) (*entities.Atendance, error)
	FindByID(id string) (*entities.Atendance, error)
}

type attendaceRepository struct {
	db *gorm.DB
}

func NewAttedanceRepository(db *gorm.DB) AttedanceRepository {
	return &attendaceRepository{db: db}
}

func (r *attendaceRepository) CreateAttedance(attedance *entities.Atendance) (*entities.Atendance, error) {
	if err := r.db.Create(attedance).Error; err != nil {
		return nil, err
	}
	return attedance, nil
}

func (r *attendaceRepository) GetAttedanceBySubjectID(subjectID string) ([]entities.Atendance, error) {
	var attedances []entities.Atendance
	// preload student and subject
	if err := r.db.Preload("Student").Preload("Subject").Where("subject_id = ?", subjectID).Find(&attedances).Error; err != nil {
		return nil, err
	}

	return attedances, nil
}

func (r *attendaceRepository) GetAttedanceByClassID(classID string) ([]entities.Atendance, error) {
	var attendances []entities.Atendance
	// preload student and subject
	if err := r.db.Preload("Student").Preload("Subject").Joins("JOIN students ON students.id = atendances.student_id").Where("students.class_id = ?", classID).Find(&attendances).Error; err != nil {
		return nil, err
	}

	return attendances, nil
}

func (r *attendaceRepository) GetAttedanceBySubjectAndClassID(subjectID, classID string) ([]entities.Atendance, error) {
	var attedances []entities.Atendance
	// preload student and subject
	if err := r.db.Preload("Student").Preload("Subject").Joins("JOIN students ON students.id = attendaces.student_id").Where("attendaces.subject_id = ? AND students.class_id = ?", subjectID, classID).Find(&attedances).Error; err != nil {
		return nil, err
	}

	return attedances, nil
}

func (r *attendaceRepository) GetMyAttedance(studentID string) ([]entities.Atendance, error) {
	var attedances []entities.Atendance
	// preload student and subject
	if err := r.db.Preload("Student").Preload("Subject").Where("student_id = ?", studentID).Find(&attedances).Error; err != nil {
		return nil, err
	}

	return attedances, nil
}

func (r *attendaceRepository) UpdateAttedance(attedance *entities.Atendance) (*entities.Atendance, error) {
	// only update status
	if err := r.db.Model(&entities.Atendance{}).Where("id = ?", attedance.ID).Select("AttendanceStatus").Updates(attedance).Error; err != nil {
		return nil, err
	}

	return attedance, nil
}

func (r *attendaceRepository) FindByID(id string) (*entities.Atendance, error) {
	var attedance entities.Atendance
	if err := r.db.Where("id = ?", id).Find(&attedance).Error; err != nil {
		return nil, err
	}

	return &attedance, nil
}
