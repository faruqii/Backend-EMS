package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Insert(task *entities.Task) error
	GetTask(id string) (*entities.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Insert(task *entities.Task) error {
	if err := r.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) GetTask(id string) (*entities.Task, error) {
	var task entities.Task

	if err := r.db.
		Preload("Class").
		Preload("Subject").
		Preload("Teacher").
		Where("id = ?", id).
		First(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}
