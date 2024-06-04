package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type AssignmentRepository interface {
	Insert(assignment *entities.StudentAssignment) error
	Update(assignment *entities.StudentAssignment) error
	FindByID(id string) (*entities.StudentAssignment, error)
	FindByTaskID(taskID string) (*entities.StudentAssignment, error)
	FindByStudentID(studentID string) (*entities.StudentAssignment, error)
	FindAll(taskID string) ([]entities.StudentAssignment, error)
	FindByTaskIDAndAssignmentID(taskID string, assignmentID string) (*entities.StudentAssignment, error)
}

type assignmentRepository struct {
	db *gorm.DB
}

func NewAssignmentRepository(db *gorm.DB) AssignmentRepository {
	return &assignmentRepository{db: db}
}

func (r *assignmentRepository) Insert(assignment *entities.StudentAssignment) error {
	return r.db.Create(assignment).Error
}

func (r *assignmentRepository) Update(assignment *entities.StudentAssignment) error {
	return r.db.Save(assignment).Error
}

func (r *assignmentRepository) FindByID(id string) (*entities.StudentAssignment, error) {
	assignment := entities.StudentAssignment{}
	if err := r.db.Preload("Task").Preload("Student").Where("id = ?", id).Find(&assignment).Error; err != nil {
		return nil, err
	}

	return &assignment, nil
}

func (r *assignmentRepository) FindByTaskID(taskID string) (*entities.StudentAssignment, error) {
	assignment := entities.StudentAssignment{}
	if err := r.db.Preload("Task").Preload("Student").Where("task_id = ?", taskID).Find(&assignment).Error; err != nil {
		return nil, err
	}

	return &assignment, nil
}

func (r *assignmentRepository) FindByStudentID(studentID string) (*entities.StudentAssignment, error) {
	assignments := entities.StudentAssignment{}
	if err := r.db.Preload("Task").Preload("Student").Where("student_id =?", studentID).Find(&assignments).Error; err != nil {
		return nil, err
	}

	return &assignments, nil
}

func (r *assignmentRepository) FindAll(taskID string) ([]entities.StudentAssignment, error) {
	assignments := []entities.StudentAssignment{}
	if err := r.db.Preload("Task").Preload("Student").Where("task_id = ?", taskID).Find(&assignments).Error; err != nil {
		return nil, err
	}

	return assignments, nil
}

func (r *assignmentRepository) FindByTaskIDAndAssignmentID(taskID string, assignmentID string) (*entities.StudentAssignment, error) {
	assignment := entities.StudentAssignment{}
	if err := r.db.Preload("Task").Preload("Student").Where("task_id = ? AND id = ?", taskID, assignmentID).Find(&assignment).Error; err != nil {
		return nil, err
	}

	return &assignment, nil
}
