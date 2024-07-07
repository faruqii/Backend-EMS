package repositories

import (
	"errors"
	"fmt"

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
	CreateSubjectMatter(subjectMatter *entities.SubjectMattter) error
	GetSubjectMatterBySubjectID(subjectID string) ([]entities.SubjectMattter, error)
	GetDetailSubjectMatter(subjectMatterID string) (*entities.SubjectMattter, error)
	GetAllSubjectInClass(classID string) ([]entities.ClassSubject, error)
	GetTeachersByClassAndSubject(classID, subjectID string) ([]entities.TeacherSubject, error)
	GetSubjectsByClassPrefix(classPrefix string) ([]entities.Subject, error)
	GetClassSubjectsByPrefixAndSubject(classPrefix string, subjectID string) ([]entities.ClassSubject, error)
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

	if len(teacherSubjects) == 0 {
		return nil, errors.New("no subjects found for the given teacher ID")
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

// GetStudentsByClassAndSubjectID returns all students in a class and subject.
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

// GetWhereIamTeachTheClass returns all class subjects where a teacher is assigned.
func (r *subjectRepository) GetWhereIamTeachTheClass(teacherID string) ([]entities.ClassSubject, error) {
	var classSubjects []entities.ClassSubject
	err := r.db.Preload("Subject").Preload("Class").Preload("Teacher").
		Where("teacher_id = ?", teacherID).Find(&classSubjects).Error
	if err != nil {
		return nil, err
	}
	return classSubjects, nil
}

// CreateSubjectMatter creates a new subject matter.
func (r *subjectRepository) CreateSubjectMatter(subjectMatter *entities.SubjectMattter) error {
    fmt.Println("Repository: Entering CreateSubjectMatter")

    tx := r.db.Begin()
    fmt.Println("Repository: Starting transaction")

    if err := tx.Create(subjectMatter).Error; err != nil {
        fmt.Println("Repository: Error creating subject matter:", err)
        tx.Rollback()
        return err
    }

    for i := range subjectMatter.Content {
        subjectMatter.Content[i].SubjectMatterID = subjectMatter.ID
        fmt.Printf("Repository: Creating content %d: %+v\n", i, subjectMatter.Content[i])
        
        // Check for duplicates before inserting
        var existingContent entities.SubjectMatterContent
        if err := tx.Where("subject_matter_id = ? AND title = ?", subjectMatter.Content[i].SubjectMatterID, subjectMatter.Content[i].Title).First(&existingContent).Error; err == nil {
            fmt.Println("Repository: Duplicate content found, skipping insertion")
            continue
        }

        if err := tx.Create(&subjectMatter.Content[i]).Error; err != nil {
            fmt.Println("Repository: Error creating content:", err)
            tx.Rollback()
            return err
        }
    }

    fmt.Println("Repository: Committing transaction")
    if err := tx.Commit().Error; err != nil {
        fmt.Println("Repository: Error committing transaction:", err)
        return err
    }

    fmt.Println("Repository: Successfully created subject matter")
    return nil
}


// GetSubjectMatterBySubjectID returns all subject matters for a subject.
func (r *subjectRepository) GetSubjectMatterBySubjectID(subjectID string) ([]entities.SubjectMattter, error) {
    var subjectMatters []entities.SubjectMattter
    if err := r.db.Preload("Subject").Preload("Content").Where("subject_id = ?", subjectID).Find(&subjectMatters).Error; err != nil {
        return nil, err
    }

    return subjectMatters, nil
}


// GetDetailSubjectMatter returns the detail of a subject matter by ID.
func (r *subjectRepository) GetDetailSubjectMatter(subjectMatterID string) (*entities.SubjectMattter, error) {
	var subjectMatter entities.SubjectMattter
	if err := r.db.Preload("Subject").Preload("Content").Where("id = ?", subjectMatterID).First(&subjectMatter).Error; err != nil {
		return nil, err
	}
	return &subjectMatter, nil
}

// GetAllSubjectInClass returns all subjects assigned to a class.
func (r *subjectRepository) GetAllSubjectInClass(classID string) ([]entities.ClassSubject, error) {
	var classSubjects []entities.ClassSubject
	if err := r.db.Preload("Class").Preload("Subject").Preload("Teacher").Where("class_id = ?", classID).Find(&classSubjects).Error; err != nil {
		return nil, err
	}
	return classSubjects, nil
}

func (r *subjectRepository) GetTeachersByClassAndSubject(classID, subjectID string) ([]entities.TeacherSubject, error) {
	var classSubjects []entities.ClassSubject
	if err := r.db.Preload("Teacher").Preload("Subject").
		Where("class_id = ? AND subject_id = ?", classID, subjectID).Find(&classSubjects).Error; err != nil {
		return nil, err
	}

	var teacherSubjects []entities.TeacherSubject
	for _, cs := range classSubjects {
		teacherSubjects = append(teacherSubjects, entities.TeacherSubject{
			Teacher: cs.Teacher,
			Subject: cs.Subject,
		})
	}

	return teacherSubjects, nil
}

func (r *subjectRepository) GetSubjectsByClassPrefix(classPrefix string) ([]entities.Subject, error) {
	var classSubjects []entities.ClassSubject
	if err := r.db.Preload("Subject").
		Joins("JOIN classes ON class_subjects.class_id = classes.id").
		Where("classes.name LIKE ?", classPrefix+"%").Find(&classSubjects).Error; err != nil {
		return nil, err
	}

	var subjects []entities.Subject
	for _, cs := range classSubjects {
		subjects = append(subjects, cs.Subject)
	}
	return subjects, nil
}

// ClassSubjectRepository.go
func (r *subjectRepository) GetClassSubjectsByPrefixAndSubject(classPrefix string, subjectID string) ([]entities.ClassSubject, error) {
	var classSubjects []entities.ClassSubject
	query := r.db.Joins("JOIN classes ON class_subjects.class_id = classes.id").
		Where("classes.name LIKE ?", classPrefix+"%")
	if subjectID != "" {
		query = query.Where("class_subjects.subject_id = ?", subjectID)
	}
	if err := query.Preload("Class").Preload("Subject").Preload("Teacher").Find(&classSubjects).Error; err != nil {
		return nil, err
	}
	return classSubjects, nil
}
