package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type QuizRepository interface {
	Insert(quiz *entities.Quiz) error
	CreateQuestion(question []entities.Question) error
	GetQuiz(id string) (*entities.Quiz, error)
	GetQuizByClassID(classID string) (*entities.Quiz, error)
	GetQuestion(quizID string) (*entities.Question, error)
}

type quizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) *quizRepository {
	return &quizRepository{db: db}
}

func (r *quizRepository) Insert(quiz *entities.Quiz) error {
	if err := r.db.Create(quiz).Error; err != nil {
		return err
	}
	return nil
}

func (r *quizRepository) CreateQuestion(question []entities.Question) error {
	// we can create more than one question
	if err := r.db.Create(question).Error; err != nil {
		return err
	}

	return nil
}

func (r *quizRepository) GetQuiz(id string) (*entities.Quiz, error) {
	var quiz entities.Quiz
	if err := r.db.First(&quiz, id).Error; err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (r *quizRepository) GetQuizByClassID(classID string) (*entities.Quiz, error) {
	var quiz entities.Quiz
	if err := r.db.Where("class_id = ?", classID).First(&quiz).Error; err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (r *quizRepository) GetQuestion(quizID string) (*entities.Question, error) {
	var question entities.Question
	if err := r.db.Where("quiz_id =?", quizID).First(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}
