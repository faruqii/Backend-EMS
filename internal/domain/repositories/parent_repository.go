package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type ParentRepository interface {
	Insert(parent *entities.Parent) error
	Update(parent *entities.Parent) error
	Delete(parent *entities.Parent) error
	FindById(id string) (*entities.Parent, error)
	InsertParentToStudent(parstud *entities.ParentStudent) error
	FindByParentAndStudent(parentID, studentID string) (*entities.ParentStudent, error)
}

type parentRepository struct {
	db *gorm.DB
}

func NewParentRepository(db *gorm.DB) *parentRepository {
	return &parentRepository{db: db}
}

func (r *parentRepository) Insert(parent *entities.Parent) error {
	if err := r.db.Create(parent).Error; err != nil {
		return err
	}
	return nil
}

func (r *parentRepository) Update(parent *entities.Parent) error {
	if err := r.db.Save(parent).Error; err != nil {
		return err
	}
	return nil
}

func (r *parentRepository) Delete(parent *entities.Parent) error {
	if err := r.db.Delete(parent).Error; err != nil {
		return err
	}
	return nil
}

func (r *parentRepository) FindById(id string) (*entities.Parent, error) {
	parent := new(entities.Parent)
	if err := r.db.Where("id = ?", id).First(parent).Error; err != nil {
		return nil, err
	}
	return parent, nil
}

func (r *parentRepository) InsertParentToStudent(parstud *entities.ParentStudent) error {
	if err := r.db.Create(parstud).Error; err != nil {
		return err
	}
	return nil
}

func (r *parentRepository) FindByParentAndStudent(parentID, studentID string) (*entities.ParentStudent, error) {
	parstud := new(entities.ParentStudent)
	if err := r.db.Where("parent_id = ? AND student_id = ?", parentID, studentID).First(parstud).Error; err != nil {
		return nil, err
	}
	return parstud, nil
}
