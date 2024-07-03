package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
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
	GetStudentIDByParentID(parentID string) (string, error)
	GetAll() ([]dto.ParentResponse, error)
	GetMyStudents(parentID string) ([]entities.ParentStudent, error)
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

func (r *parentRepository) GetStudentIDByParentID(parentID string) (string, error) {
	parstud := new(entities.ParentStudent)
	if err := r.db.Where("parent_id = ?", parentID).First(parstud).Error; err != nil {
		return "", err
	}
	return parstud.StudentID, nil
}

func (r *parentRepository) GetAll() ([]dto.ParentResponse, error) {
	var parents []dto.ParentResponse
	// join table with ParentStudent to show the child of parent
	err := r.db.Table("parents").
		Select("parents.id, parents.name, parents.address, parents.occupation, parents.phone_number, parents.email, students.name as student_name").
		Joins("LEFT JOIN parent_students ON parents.id = parent_students.parent_id").
		Joins("LEFT JOIN students ON parent_students.student_id = students.id").
		Scan(&parents).Error
	if err != nil {
		return nil, err
	}

	return parents, nil
}

func (r *parentRepository) GetMyStudents(parentID string) ([]entities.ParentStudent, error) {
	// preloading the student
	var parstuds []entities.ParentStudent
	if err := r.db.Preload("Student").Where("parent_id = ?", parentID).Find(&parstuds).Error; err != nil {
		return nil, err
	}
	return parstuds, nil
}
