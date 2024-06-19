package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

// StudentRepository is a contract of student repository
// This is an interface that will be implemented by StudentRepository

type StudentRepository interface {
	Insert(student *entities.Student) error
	Update(student *entities.Student) error
	Delete(student *entities.Student) error
	FindById(id string) (*entities.Student, error)
	FindByNISN(nisn string) (*entities.Student, error)
	FindStudentByToken(token string) (string, error)
	FindRoleByName(name string) (*entities.Role, error)
	FindByUsername(username string) (*entities.User, error)
	InsertStudentToClass(studentID, classID string) (*entities.Student, error)
	GetAllStudents() ([]entities.Student, error)
	IsStudentAlreadyInClass(studentID string) (bool, error)
	GetAllStudentsByClassID(classID string) ([]entities.Student, error)
	FindStudentClassIDByStudentID(studentID string) (string, error)
	GetStudentByUserID(userID string) (*entities.Student, error)
	RemoveStudentFromClass(studentID string) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *studentRepository {
	return &studentRepository{db: db}
}

func (r *studentRepository) Insert(student *entities.Student) error {
	if err := r.db.Create(student).Error; err != nil {
		return err
	}
	return nil
}

func (r *studentRepository) Update(student *entities.Student) error {
	if err := r.db.Save(student).Error; err != nil {
		return err
	}
	return nil
}

func (r *studentRepository) Delete(student *entities.Student) error {
	if err := r.db.Delete(student).Error; err != nil {
		return err
	}
	return nil
}

func (r *studentRepository) FindById(id string) (*entities.Student, error) {
	student := new(entities.Student)
	if err := r.db.Where("id = ?", id).First(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (r *studentRepository) FindByNISN(nisn string) (*entities.Student, error) {
	student := new(entities.Student)
	if err := r.db.Where("nisn = ?", nisn).First(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (r *studentRepository) FindStudentByToken(token string) (string, error) {
	var studentID string
	if err := r.db.Raw("SELECT user_id FROM tokens WHERE token = ?", token).Scan(&studentID).Error; err != nil {
		return "", err
	}
	return studentID, nil
}

func (r *studentRepository) FindRoleByName(name string) (*entities.Role, error) {
	role := new(entities.Role)
	if err := r.db.Where("name = ?", name).First(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *studentRepository) FindByUsername(username string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *studentRepository) InsertStudentToClass(studentID, classID string) (*entities.Student, error) {
	// Fetch the student along with the associated class
	var student entities.Student
	if err := r.db.Preload("Class").Where("id = ?", studentID).First(&student).Error; err != nil {
		return nil, err
	}

	// Update the class ID
	student.ClassID = &classID
	if err := r.db.Save(&student).Error; err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *studentRepository) GetAllStudents() ([]entities.Student, error) {
	var students []entities.Student
	if err := r.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *studentRepository) IsStudentAlreadyInClass(studentID string) (bool, error) {
	var student entities.Student
	if err := r.db.Where("id = ?", studentID).First(&student).Error; err != nil {
		return false, err
	}
	if student.ClassID == nil {
		return false, nil
	}
	return true, nil
}

func (r *studentRepository) GetAllStudentsByClassID(classID string) ([]entities.Student, error) {
	var students []entities.Student
	if err := r.db.Where("class_id = ?", classID).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *studentRepository) FindStudentClassIDByStudentID(studentID string) (string, error) {
	var classID string
	if err := r.db.Raw("SELECT class_id FROM students WHERE id =?", studentID).Scan(&classID).Error; err != nil {
		return "", err
	}
	return classID, nil
}

func (r *studentRepository) GetStudentByUserID(userID string) (*entities.Student, error) {
	var student entities.Student
	if err := r.db.Where("user_id = ?", userID).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepository) RemoveStudentFromClass(studentID string) error {
    var student entities.Student
    if err := r.db.Where("id = ?", studentID).First(&student).Error; err != nil {
        return err
    }

    // Ensuring ClassID is nil
    student.ClassID = nil

    // Updating the ClassID field specifically
    if err := r.db.Model(&student).Update("class_id", gorm.Expr("NULL")).Error; err != nil {
        return err
    }

    return nil
}


// Path: internal/domain/repositories/student_repository.go
