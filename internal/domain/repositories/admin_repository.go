package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByUsername(username string) (*entities.Admin, error)
	FindById(id string) (*entities.Admin, error)
	FindAdminByToken(token string) (string, error)
	CreateClass(class *entities.Class) (*entities.Class, error)
	UpdateClass(class *entities.Class) (*entities.Class, error)
	DeleteClass(class *entities.Class) error
	FindClassById(id int) (*entities.Class, error)
	GetAllClass() ([]entities.Class, error)
	CreateSubject(subject *entities.Subject) (*entities.Subject, error)
	UpdateSubject(subject *entities.Subject) (*entities.Subject, error)
	DeleteSubject(subject *entities.Subject) error
	FindSubjectById(id int) (*entities.Subject, error)
	GetAllSubject() ([]entities.Subject, error)
	CreateTeacher(teacher *entities.Teacher) (*entities.Teacher, error)
	UpdateTeacher(teacher *entities.Teacher) (*entities.Teacher, error)
	DeleteTeacher(teacher *entities.Teacher) error
	FindTeacherById(id int) (*entities.Teacher, error)
	GetAllTeacher() ([]entities.Teacher, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db: db}
}