package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Insert(task *entities.Task) error
	GetTask(id string) (*entities.Task, error)
	GetTaskByClassID(classID string) ([]entities.Task, error)
	GetTaskByTeacherID(teacherID string) ([]entities.Task, error)
	Update(taskID string, task *entities.Task) error
	Delete(taskID string) error
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

func (r *taskRepository) GetTaskByClassID(classID string) ([]entities.Task, error) {
	var task []entities.Task

	if err := r.db.
		Preload("Class").
		Preload("Subject").
		Preload("Teacher").
		Where("class_id = ?", classID).
		Find(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepository) GetTaskByTeacherID(teacherID string) ([]entities.Task, error) {
	var tasks []entities.Task

	if err := r.db.
		Preload("Class").
		Preload("Subject").
		Preload("Teacher").
		Where("teacher_id = ?", teacherID).
		Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) Update(taskID string, task *entities.Task) error {
	if err := r.db.Model(&entities.Task{}).
		Where("id = ?", taskID).
		Updates(task).Error; err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) Delete(taskID string) error {
	if err := r.db.Where("id = ?", taskID).Delete(&entities.Task{}).Error; err != nil {
		return err
	}

	return nil
}
