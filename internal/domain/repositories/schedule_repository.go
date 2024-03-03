package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	Insert(schedule *entities.Schedule) error
	Update(schedule *entities.Schedule) error
	Delete(id string) error
	FindByID(id string) (*entities.Schedule, error)
	FindByClassID(classID string) ([]entities.Schedule, error)
	GetAll() ([]entities.Schedule, error)
	GetScheduleByID(id string) (*entities.Schedule, error)
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (r *scheduleRepository) Insert(schedule *entities.Schedule) error {
	if err := r.db.Create(schedule).Error; err != nil {
		return err
	}
	return nil
}

func (r *scheduleRepository) Update(schedule *entities.Schedule) error {
	if err := r.db.Save(schedule).Error; err != nil {
		return err
	}
	return nil
}

func (r *scheduleRepository) Delete(id string) error {
	if err := r.db.Delete(&entities.Schedule{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *scheduleRepository) FindByID(id string) (*entities.Schedule, error) {
	var schedule entities.Schedule
	if err := r.db.First(&schedule, id).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *scheduleRepository) FindByClassID(classID string) ([]entities.Schedule, error) {
	var schedules []entities.Schedule
	if err := r.db.Where("class_id = ?", classID).Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *scheduleRepository) GetAll() ([]entities.Schedule, error) {
	var schedules []entities.Schedule
	// Preload the class, subject, and teacher
	if err := r.db.Preload("Class").Preload("Subject").Preload("Teacher").Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *scheduleRepository) GetScheduleByID(id string) (*entities.Schedule, error) {
	var schedule entities.Schedule
	if err := r.db.Preload("Class").Preload("Subject").Preload("Teacher").First(&schedule, id).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}


