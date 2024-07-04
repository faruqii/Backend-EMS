package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type AssignmentRepository interface {
	Insert(assignment *entities.StudentAssignment) error
	Update(assignment *entities.StudentAssignment) error
	FindByID(id string) (*entities.StudentAssignment, error)
	FindByTaskID(taskID string) (*entities.StudentAssignment, error)
	FindByStudentID(studentID string) (*entities.StudentAssignment, error)
	FindAll(taskID string) ([]entities.StudentAssignment, error)
	FindByTaskIDAndAssignmentID(taskID string, assignmentID string) (*entities.StudentAssignment, error)
	InsertQuiz(assignment *entities.StudentQuizAssignment) error
	FindByQuizID(quizID string) (*entities.StudentQuizAssignment, error)
	UpdateQuizAssignment(assignment *entities.StudentQuizAssignment) error
	GetAllQuizAssignment(quizID string) ([]entities.StudentQuizAssignment, error)
	GetQuizAssignment(quizAssignmentID string) (*entities.StudentQuizAssignment, error)
	GetQuizByStudentID(studentID string) ([]entities.StudentQuizAssignment, error)
	GetStudentQuizAssignment(quizID, studentID string) (*entities.StudentQuizAssignment, error)
	GradeStudentQuiz(quizAssignmentID string, status string, grade float64) error
	GetMyQuizAssignment(studentID string, subjectID string) ([]entities.StudentQuizAssignment, error)
}

type assignmentRepository struct {
	db *gorm.DB
}

func NewAssignmentRepository(db *gorm.DB) AssignmentRepository {
	return &assignmentRepository{db: db}
}

func (r *assignmentRepository) Insert(assignment *entities.StudentAssignment) error {
	return r.db.Create(assignment).Error
}

func (r *assignmentRepository) Update(assignment *entities.StudentAssignment) error {
	return r.db.Model(&entities.StudentAssignment{}).
		Where("id = ?", assignment.ID).
		Select("Grade", "Feedback"). // Only update these fields
		Updates(assignment).Error
}

func (r *assignmentRepository) FindByID(id string) (*entities.StudentAssignment, error) {
	assignment := entities.StudentAssignment{}
	if err := r.db.Preload("Task").Preload("Student").Where("id = ?", id).Find(&assignment).Error; err != nil {
		return nil, err
	}

	return &assignment, nil
}

func (r *assignmentRepository) FindByTaskID(taskID string) (*entities.StudentAssignment, error) {
	assignment := entities.StudentAssignment{}
	if err := r.db.Preload("Task").Preload("Student").Where("task_id = ?", taskID).Find(&assignment).Error; err != nil {
		return nil, err
	}

	return &assignment, nil
}

func (r *assignmentRepository) FindByStudentID(studentID string) (*entities.StudentAssignment, error) {
	assignments := entities.StudentAssignment{}
	if err := r.db.Preload("Task").Preload("Student").Where("student_id =?", studentID).Find(&assignments).Error; err != nil {
		return nil, err
	}

	return &assignments, nil
}

func (r *assignmentRepository) FindAll(taskID string) ([]entities.StudentAssignment, error) {
	assignments := []entities.StudentAssignment{}
	if err := r.db.Preload("Task").Preload("Student").Where("task_id = ?", taskID).Find(&assignments).Error; err != nil {
		return nil, err
	}

	return assignments, nil
}

func (r *assignmentRepository) FindByTaskIDAndAssignmentID(taskID string, assignmentID string) (*entities.StudentAssignment, error) {
	assignment := entities.StudentAssignment{}
	if err := r.db.Preload("Task").Preload("Student").Where("task_id = ? AND id = ?", taskID, assignmentID).Find(&assignment).Error; err != nil {
		return nil, err
	}

	return &assignment, nil
}

func (r *assignmentRepository) InsertQuiz(assignment *entities.StudentQuizAssignment) error {
	return r.db.Create(assignment).Error
}

func (r *assignmentRepository) FindByQuizID(id string) (*entities.StudentQuizAssignment, error) {
	var assignment entities.StudentQuizAssignment
	if err := r.db.Preload("Quiz.Questions").First(&assignment, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (r *assignmentRepository) UpdateQuizAssignment(assignment *entities.StudentQuizAssignment) error {
	return r.db.Save(assignment).Error
}

func (r *assignmentRepository) GetAllQuizAssignment(quizID string) ([]entities.StudentQuizAssignment, error) {
	var assignments []entities.StudentQuizAssignment
	if err := r.db.Preload("Quiz").Preload("Student").Where("quiz_id = ?", quizID).Find(&assignments).Error; err != nil {
		return nil, err
	}
	return assignments, nil
}

func (r *assignmentRepository) GetQuizAssignment(quizAssignmentID string) (*entities.StudentQuizAssignment, error) {
	var assignment entities.StudentQuizAssignment
	if err := r.db.Preload("Quiz.Questions").First(&assignment, "id = ?", quizAssignmentID).Error; err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (r *assignmentRepository) GetQuizByStudentID(studentID string) ([]entities.StudentQuizAssignment, error) {
	var assignments []entities.StudentQuizAssignment
	if err := r.db.Preload("Quiz").Preload("Student").Where("student_id = ?", studentID).Find(&assignments).Error; err != nil {
		return nil, err
	}
	return assignments, nil
}

func (r *assignmentRepository) GetStudentQuizAssignment(quizID, studentID string) (*entities.StudentQuizAssignment, error) {
	var assignment entities.StudentQuizAssignment
	if err := r.db.Preload("Quiz").Preload("Student").First(&assignment, "quiz_id = ? AND student_id = ?", quizID, studentID).Error; err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (r *assignmentRepository) GradeStudentQuiz(quizAssignmentID string, status string, grade float64) error {
	return r.db.Model(&entities.StudentQuizAssignment{}).
		Where("id = ?", quizAssignmentID).
		Select("Status", "Grade"). // Only update these fields
		Updates(map[string]interface{}{"Status": status, "Grade": grade}).Error
}

func (r *assignmentRepository) GetMyQuizAssignment(studentID string, subjectID string) ([]entities.StudentQuizAssignment, error) {
	var assignments []entities.StudentQuizAssignment
	db := r.db.Model(&entities.StudentQuizAssignment{}).
		Where("student_quiz_assignments.student_id = ?", studentID)

	if subjectID != "" {
		db = db.Joins("JOIN quizzes ON quizzes.id = student_quiz_assignments.quiz_id").
			Where("quizzes.subject_id = ?", subjectID)
	}

	// Apply the preloads and execute the query
	err := db.Preload("Quiz").Preload("Student").Preload("Quiz.Subject").Find(&assignments).Error
	if err != nil {
		return nil, err
	}

	return assignments, nil
}
