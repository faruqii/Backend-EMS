package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type GradeRepository interface {
	InsertGrade(grade *entities.Grade) (*entities.Grade, error)
	GetGradeByID(gradeID string) (*entities.Grade, error)
	GetAllGradeByStudentID(studentID string) ([]entities.Grade, error)
	GetAllGradeBySubjectID(subjectID string) ([]entities.Grade, error)
	GetAllGrade() ([]entities.Grade, error)
	FilterBySemester(semester string) ([]entities.Grade, error)
	FilterByAcademicYear(academicYear string) ([]entities.Grade, error)
	FilterBySemesterAndAcademicYear(semester, academicYear string) ([]entities.Grade, error)
	FilterStudentGradesBySemester(studentID, semester string) ([]entities.Grade, error)
	FilterStudentGradesByAcademicYear(studentID, academicYear string) ([]entities.Grade, error)
	FilterStudentGradesBySemesterAndAcademicYear(studentID, semester, academicYear string) ([]entities.Grade, error)
	UpdateGrade(grade *entities.Grade) (*entities.Grade, error)
}

type gradeRepository struct {
	db *gorm.DB
}

func NewGradeRepository(db *gorm.DB) GradeRepository {
	return &gradeRepository{db: db}
}

func (r *gradeRepository) InsertGrade(grade *entities.Grade) (*entities.Grade, error) {
	if err := r.db.Create(grade).Error; err != nil {
		return nil, err
	}

	return grade, nil
}

func (r *gradeRepository) GetGradeByID(gradeID string) (*entities.Grade, error) {
	// preload
	var grade entities.Grade
	if err := r.db.Preload("Student").Preload("Subject").Preload("Teacher").Where("id = ?", gradeID).First(&grade).Error; err != nil {
		return nil, err
	}

	return &grade, nil
}

func (r *gradeRepository) GetAllGradeByStudentID(studentID string) ([]entities.Grade, error) {
	// preload
	var grades []entities.Grade
	if err := r.db.Preload("Student").Preload("Subject").Preload("Teacher").Where("student_id = ?", studentID).Find(&grades).Error; err != nil {
		return nil, err
	}

	return grades, nil
}

func (r *gradeRepository) GetAllGradeBySubjectID(subjectID string) ([]entities.Grade, error) {
	// preload
	var grades []entities.Grade
	if err := r.db.Preload("Student").Preload("Subject").Preload("Teacher").Where("subject_id = ?", subjectID).Find(&grades).Error; err != nil {
		return nil, err
	}

	return grades, nil
}

func (r *gradeRepository) GetAllGrade() ([]entities.Grade, error) {
	// preload
	var grades []entities.Grade
	if err := r.db.Preload("Student").Preload("Subject").Preload("Teacher").Find(&grades).Error; err != nil {
		return nil, err
	}

	return grades, nil
}

func (r *gradeRepository) FilterBySemester(semester string) ([]entities.Grade, error) {
	// preload
	var grades []entities.Grade
	if err := r.db.Preload("Student").Preload("Subject").Preload("Teacher").Where("semester = ?", semester).Find(&grades).Error; err != nil {
		return nil, err
	}

	return grades, nil
}

func (r *gradeRepository) FilterByAcademicYear(academicYear string) ([]entities.Grade, error) {
	// preload
	var grades []entities.Grade
	if err := r.db.Preload("Student").Preload("Subject").Preload("Teacher").Where("academic_year = ?", academicYear).Find(&grades).Error; err != nil {
		return nil, err
	}

	return grades, nil
}

func (r *gradeRepository) FilterBySemesterAndAcademicYear(semester, academicYear string) ([]entities.Grade, error) {
	var grades []entities.Grade
	if err := r.db.Preload("Student").Preload("Subject").Preload("Teacher").
		Where("semester = ? AND academic_year = ?", semester, academicYear).Find(&grades).Error; err != nil {
		return nil, err
	}
	return grades, nil
}

func (r *gradeRepository) FilterStudentGradesBySemester(studentID, semester string) ([]entities.Grade, error) {
	var grades []entities.Grade
	if err := r.db.Preload("Student").Preload("Subject").Preload("Teacher").
		Where("student_id = ? AND semester = ?", studentID, semester).Find(&grades).Error; err != nil {
		return nil, err
	}

	return grades, nil
}

func (r *gradeRepository) FilterStudentGradesByAcademicYear(studentID, academicYear string) ([]entities.Grade, error) {
	var grades []entities.Grade
	if err := r.db.Preload("Student").Preload("Subject").Preload("Teacher").
		Where("student_id = ? AND academic_year = ?", studentID, academicYear).Find(&grades).Error; err != nil {
		return nil, err
	}

	return grades, nil
}

func (r *gradeRepository) FilterStudentGradesBySemesterAndAcademicYear(studentID, semester, academicYear string) ([]entities.Grade, error) {
	var grades []entities.Grade
	if err := r.db.Preload("Student").Preload("Subject").Preload("Teacher").
		Where("student_id = ? AND semester = ? AND academic_year = ?", studentID, semester, academicYear).Find(&grades).Error; err != nil {
		return nil, err
	}

	return grades, nil
}

func (r *gradeRepository) UpdateGrade(grade *entities.Grade) (*entities.Grade, error) {
	// find grade
	var oldGrade entities.Grade
	if err := r.db.Where("id = ?", grade.ID).First(&oldGrade).Error; err != nil {
		return nil, err
	}

	// update grade
	if err := r.db.Model(&oldGrade).Updates(grade).Error; err != nil {
		return nil, err
	}

	return grade, nil
}
