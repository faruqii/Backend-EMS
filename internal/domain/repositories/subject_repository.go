package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

// SubjectRepository defines methods to interact with the Subject model.
type SubjectRepository interface {
	FindByID(id string) (*entities.Subject, error)
	Create(subject *entities.Subject) error
	Update(subject *entities.Subject) error
	Delete(id string) error
	GetAll() ([]entities.Subject, error)
	AssignTeacherToSubject(teacherID, subjectID string) error
	IsTeacherAssignedToSubject(teacherID, subjectID string) (bool, error)
	GetTeachersBySubjectID(subjectID string) ([]entities.TeacherSubject, error)
	GetTeacherSubjects(teacherID string) ([]entities.TeacherSubject, error)
	AssignSubjectToClass(subjectID, teacherID, classID string) (*entities.ClassSubject, error)
	GetClassSubjects(classID string) ([]entities.ClassSubject, error)
	GetStudentsByClassAndSubjectID(classID, subjectID string) ([]entities.Student, error)
	GetWhereIamTeachTheClass(teacherID string) ([]entities.ClassSubject, error)
}

// subjectRepository is a concrete implementation of SubjectRepository.
type subjectRepository struct {
	db *gorm.DB
}

// NewSubjectRepository creates a new instance of subjectRepository.
func NewSubjectRepository(db *gorm.DB) SubjectRepository {
	return &subjectRepository{db: db}
}

// FindByID finds a subject by ID.
func (r *subjectRepository) FindByID(id string) (*entities.Subject, error) {
	var subject entities.Subject
	if err := r.db.Where("id = ?", id).First(&subject).Error; err != nil {
		return nil, err
	}
	return &subject, nil
}

// Create creates a new subject.
func (r *subjectRepository) Create(subject *entities.Subject) error {
	if err := r.db.Create(subject).Error; err != nil {
		return err
	}
	return nil
}

// Update updates an existing subject.
func (r *subjectRepository) Update(subject *entities.Subject) error {
	if err := r.db.Save(subject).Error; err != nil {
		return err
	}
	return nil
}

// Delete deletes a subject by ID.
func (r *subjectRepository) Delete(id string) error {
	if err := r.db.Delete(&entities.Subject{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetAll returns all subjects.
func (r *subjectRepository) GetAll() ([]entities.Subject, error) {
	var subjects []entities.Subject
	if err := r.db.Find(&subjects).Error; err != nil {
		return nil, err
	}
	return subjects, nil
}

// AssignTeacherToSubject assigns a teacher to a subject.
func (r *subjectRepository) AssignTeacherToSubject(teacherID, subjectID string) error {
	teacherSubject := entities.TeacherSubject{
		TeacherID: teacherID,
		SubjectID: subjectID,
	}
	if err := r.db.Create(&teacherSubject).Error; err != nil {
		return err
	}
	return nil
}

// IsTeacherAssignedToSubject checks if a teacher is assigned to a subject.
func (r *subjectRepository) IsTeacherAssignedToSubject(teacherID, subjectID string) (bool, error) {
	var count int64
	if err := r.db.Model(&entities.TeacherSubject{}).
		Where("teacher_id = ? AND subject_id = ?", teacherID, subjectID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetTeachersBySubjectID returns all teachers assigned to a subject.
func (r *subjectRepository) GetTeachersBySubjectID(subjectID string) ([]entities.TeacherSubject, error) {
	var teacherSubjects []entities.TeacherSubject
	if err := r.db.Preload("Teacher").Preload("Subject").
		Where("subject_id = ?", subjectID).Find(&teacherSubjects).Error; err != nil {
		return nil, err
	}
	return teacherSubjects, nil
}

// GetTeacherSubjects returns all subjects assigned to a teacher.
func (r *subjectRepository) GetTeacherSubjects(teacherID string) ([]entities.TeacherSubject, error) {
	var teacherSubjects []entities.TeacherSubject
	if err := r.db.Preload("Teacher").Preload("Subject").
		Where("teacher_id = ?", teacherID).Find(&teacherSubjects).Error; err != nil {
		return nil, err
	}
	return teacherSubjects, nil
}

// AssignSubjectToClass assigns a subject to a class.
func (r *subjectRepository) AssignSubjectToClass(subjectID, teacherID, classID string) (*entities.ClassSubject, error) {
	classSubject := entities.ClassSubject{
		ClassID:   classID,
		SubjectID: subjectID,
		TeacherID: teacherID,
	}
	if err := r.db.Create(&classSubject).Error; err != nil {
		return nil, err
	}

	return &classSubject, nil
}

// GetClassSubjects returns all subjects assigned to a class.
func (r *subjectRepository) GetClassSubjects(classID string) ([]entities.ClassSubject, error) {
	var classSubjects []entities.ClassSubject
	if err := r.db.Preload("Subject").Preload("Teacher").Preload("Class").
		Where("class_id = ?", classID).Find(&classSubjects).Error; err != nil {
		return nil, err
	}
	return classSubjects, nil
}

func (r *subjectRepository) GetStudentsByClassAndSubjectID(classID, subjectID string) ([]entities.Student, error) {
	var students []entities.Student
	err := r.db.Model(&entities.Student{}).
		Joins("JOIN classes ON students.class_id = classes.id").
		Joins("JOIN class_subjects ON classes.id = class_subjects.class_id").
		Where("class_subjects.class_id = ? AND class_subjects.subject_id = ?", classID, subjectID).
		Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (r *subjectRepository) GetWhereIamTeachTheClass(teacherID string) ([]entities.ClassSubject, error) {
	var classSubjects []entities.ClassSubject
	err := r.db.Preload("Subject").Preload("Class").Preload("Teacher").
		Where("teacher_id = ?", teacherID).Find(&classSubjects).Error
	if err != nil {
		return nil, err
	}
	return classSubjects, nil
}
