package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type ClassRepository interface {
	Insert(class *entities.Class) error
	Update(class *entities.Class) error
	Delete(id string) error
	FindByID(id string) (*entities.Class, error)
	FindByTeacherID(teacherID string) ([]entities.Class, error)
	GetAll() ([]entities.Class, error)
	FindByName(name string) (*entities.Class, error)
	GetAllStudents(classID string) ([]entities.Student, error)
	ClassExists(classID string) (bool, error)
	IsTeacherTeachTheClass(classID string) (bool, error)
	IsTeacherHomeRoomTeacher(teacherID, classID string) (bool, error)
	RemoveStudentsFromClass(classID string) error
	GetClassesByPrefix(name string) ([]entities.Class, error)
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{db: db}
}

func (r *classRepository) Insert(class *entities.Class) error {
	if err := r.db.Create(class).Error; err != nil {
		return err
	}
	return nil
}

func (r *classRepository) Update(class *entities.Class) error {
	if err := r.db.Save(class).Error; err != nil {
		return err
	}
	return nil
}

func (r *classRepository) Delete(id string) error {
	if err := r.db.Delete(&entities.Class{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *classRepository) FindByID(id string) (*entities.Class, error) {
	var class entities.Class
	if err := r.db.Preload("HomeRoomTeacher").First(&class, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

func (r *classRepository) FindByTeacherID(teacherID string) ([]entities.Class, error) {
	var classes []entities.Class
	if err := r.db.Where("teacher_id = ?", teacherID).Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

func (r *classRepository) GetAll() ([]entities.Class, error) {
	var classes []entities.Class
	// preloading the teacher
	if err := r.db.Preload("HomeRoomTeacher").Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

func (r *classRepository) FindByName(name string) (*entities.Class, error) {
	var class entities.Class
	if err := r.db.First(&class, "name = ?", name).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

func (r *classRepository) GetAllStudents(classID string) ([]entities.Student, error) {
	var students []entities.Student
	if err := r.db.Where("class_id = ?", classID).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *classRepository) ClassExists(classID string) (bool, error) {
	var count int64
	if err := r.db.Model(&entities.Class{}).Where("id = ?", classID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil

}

func (r *classRepository) IsTeacherTeachTheClass(classID string) (bool, error) {
	var count int64
	if err := r.db.Model(&entities.Class{}).Where("id =?", classID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *classRepository) IsTeacherHomeRoomTeacher(teacherID, classID string) (bool, error) {
	var count int64
	if err := r.db.Model(&entities.Class{}).Where("home_room_teacher_id = ? AND id = ?", teacherID, classID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *classRepository) RemoveStudentsFromClass(classID string) error {
	if err := r.db.Model(&entities.Student{}).Where("class_id = ?", classID).Update("class_id", nil).Error; err != nil {
		return err
	}
	return nil
}

func (r *classRepository) GetClassesByPrefix(name string) ([]entities.Class, error) {
	var classes []entities.Class
	if err := r.db.Preload("HomeRoomTeacher").Where("name LIKE ?", name+"%").Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}