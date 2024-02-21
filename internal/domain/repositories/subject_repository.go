package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

// SubjectRepository defines methods to interact with the Subject model.
type SubjectRepository interface {
	FindByID(id string) (*entities.Subject, error)
	Create(subject *entities.Subject) error
	Update(subject *entities.Subject) error
	Delete(id string) error
}

// subjectRepository is a concrete implementation of SubjectRepository.
type subjectRepository struct {
	db *gorm.DB
}

// NewSubjectRepository creates a new instance of subjectRepository.
func NewSubjectRepository(db *gorm.DB) SubjectRepository {
	return &subjectRepository{db: db}
}

// FindByID finds a subject by ID.
func (r *subjectRepository) FindByID(id string) (*entities.Subject, error) {
	var subject entities.Subject
	if err := r.db.First(&subject, id).Error; err != nil {
		return nil, err
	}
	return &subject, nil
}

// Create creates a new subject.
func (r *subjectRepository) Create(subject *entities.Subject) error {
	if err := r.db.Create(subject).Error; err != nil {
		return err
	}
	return nil
}

// Update updates an existing subject.
func (r *subjectRepository) Update(subject *entities.Subject) error {
	if err := r.db.Save(subject).Error; err != nil {
		return err
	}
	return nil
}

// Delete deletes a subject by ID.
func (r *subjectRepository) Delete(id string) error {
	if err := r.db.Delete(&entities.Subject{}, id).Error; err != nil {
		return err
	}
	return nil
}
