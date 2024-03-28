package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GradeRepository interface {
	Insert(grade *entities.Grade) error
	Update(grade *entities.Grade) error
	Delete(grade *entities.Grade) error
	FindByID(gradeID uuid.UUID) (*entities.Grade, error)
	FindByStudentID(studentID string) ([]entities.Grade, error)
	FindBySubjectID(subjectID string) ([]entities.Grade, error)
	FindByTeacherID(teacherID string) ([]entities.Grade, error)
	FindAll() ([]entities.Grade, error)
}

type gradeRepository struct {
	db *gorm.DB
}

func NewGradeRepository(db *gorm.DB) *gradeRepository {
	return &gradeRepository{db: db}
}

func (r *gradeRepository) Insert(grade *entities.Grade) error {
	return r.db.Create(grade).Error
}

func (r *gradeRepository) Update(grade *entities.Grade) error {
	return r.db.Save(grade).Error
}

func (r *gradeRepository) Delete(grade *entities.Grade) error {
	return r.db.Delete(grade).Error
}

func (r *gradeRepository) FindByID(gradeID uuid.UUID) (*entities.Grade, error) {
	var grade entities.Grade
	if err := r.db.Preload("Student").Preload("Subject").Preload("Teacher").Where("id = ?", gradeID).First(&grade).Error; err != nil {
		return nil, err
	}
	return &grade, nil
}

func (r *gradeRepository) FindByStudentID(studentID string) ([]entities.Grade, error) {
	var grades []entities.Grade
	if err := r.db.Preload("Student").Preload("Subject").Preload("Teacher").Where("student_id = ?", studentID).Find(&grades).Error; err != nil {
		return nil, err
	}
	return grades, nil
}

func (r *gradeRepository) FindBySubjectID(subjectID string) ([]entities.Grade, error) {
	var grades []entities.Grade
	if err := r.db.Preload("Student").Preload("Subject").Preload("Teacher").Where("subject_id = ?", subjectID).Find(&grades).Error; err != nil {
		return nil, err
	}
	return grades, nil
}

func (r *gradeRepository) FindByTeacherID(teacherID string) ([]entities.Grade, error) {
	var grades []entities.Grade
	if err := r.db.Preload("Student").Preload("Subject").Preload("Teacher").Where("teacher_id = ?", teacherID).Find(&grades).Error; err != nil {
		return nil, err
	}
	return grades, nil
}

func (r *gradeRepository) FindAll() ([]entities.Grade, error) {
	var grades []entities.Grade
	if err := r.db.Preload("Student").Preload("Subject").Preload("Teacher").Find(&grades).Error; err != nil {
		return nil, err
	}
	return grades, nil
}