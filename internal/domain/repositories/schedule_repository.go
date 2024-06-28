package repositories

import (
	"time"

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
	GetPreloadSchedule() (*entities.Schedule, error)
	IsScheduleExists(classID, subjectID string) (bool, error)
	GetTeacherTodaySchedule(teacherID string, dayOfWeek time.Weekday) ([]entities.Schedule, error)
	GetAllTeacherSchedule(teacherID string) ([]entities.Schedule, error)
	GetStudentTodaySchedule(studentID string, dayOfWeek time.Weekday) ([]entities.Schedule, error)
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
	// dont use save because it will update all fields and also it related entities
	if err := r.db.Model(&entities.Schedule{}).Where("id = ?", schedule.ID).Updates(schedule).Error; err != nil {
		return err
	}
	return nil
}

func (r *scheduleRepository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.Schedule{}).Error; err != nil {
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
	// Preload the class, subject, and teacher
	if err := r.db.Preload("Class").
		Preload("Subject").
		Preload("Teacher").
		Where("class_id = ?", classID).
		Find(&schedules).Error; err != nil {
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

func (r *scheduleRepository) GetPreloadSchedule() (*entities.Schedule, error) {
	var schedules entities.Schedule
	if err := r.db.Preload("Class").Preload("Subject").Preload("Teacher").Find(&schedules).Error; err != nil {
		return nil, err
	}
	return &schedules, nil
}

func (r *scheduleRepository) IsScheduleExists(classID, subjectID string) (bool, error) {
	var count int64
	if err := r.db.Model(&entities.Schedule{}).
		Where("class_id = ? AND subject_id = ?", classID, subjectID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *scheduleRepository) GetTeacherTodaySchedule(teacherID string, dayOfWeek time.Weekday) ([]entities.Schedule, error) {
	var schedules []entities.Schedule
	// Preload the class, subject, and teacher
	if err := r.db.Preload("Class").Preload("Subject").Preload("Teacher").
		Where("teacher_id = ? AND day_of_week = ?", teacherID, dayOfWeek).
		Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *scheduleRepository) GetAllTeacherSchedule(teacherID string) ([]entities.Schedule, error) {
	var schedules []entities.Schedule
	// Preload the class, subject, and teacher
	if err := r.db.Preload("Class").Preload("Subject").Preload("Teacher").
		Where("teacher_id = ?", teacherID).
		Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *scheduleRepository) GetStudentTodaySchedule(studentID string, dayOfWeek time.Weekday) ([]entities.Schedule, error) {
	var schedules []entities.Schedule
	// Preload the class, subject, and teacher
	if err := r.db.Preload("Class").Preload("Subject").Preload("Teacher").
		Joins("JOIN class_students ON class_students.class_id = schedules.class_id").
		Where("class_students.student_id = ? AND schedules.day_of_week = ?", studentID, dayOfWeek).
		Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}
