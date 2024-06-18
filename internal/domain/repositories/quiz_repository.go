package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type QuizRepository interface {
	Insert(quiz *entities.Quiz) error
	CreateQuestion(question []entities.Question) error
	GetQuiz(id string) (*entities.Quiz, error)
	GetQuizByClassID(classID string) ([]entities.Quiz, error)
	GetQuestion(quizID string) ([]entities.Question, error)
	GetQuizBySubjectID(subjectID string) (*entities.Quiz, error)
	GetQuizByTeacherID(teacherID string) ([]entities.Quiz, error)
	CountQuestion(quizID string) (int64, error)
	MatchAnswer(quizID string, answer []string) (int64, error)
}

type quizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) *quizRepository {
	return &quizRepository{db: db}
}

func (r *quizRepository) Insert(quiz *entities.Quiz) error {
	// Insert the quiz, and handle conflict by doing nothing or updating fields
	if err := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(quiz).Error; err != nil {
		return err
	}
	return nil
}

func (r *quizRepository) CreateQuestion(questions []entities.Question) error {
	for i := range questions {
		// Check if the question already exists
		var count int64
		if err := r.db.Model(&entities.Question{}).
			Where("quiz_id = ? AND text = ?", questions[i].QuizID, questions[i].Text).
			Count(&count).Error; err != nil {
			return err
		}

		if count == 0 {
			// Convert []string to pq.StringArray
			optionsArray := pq.StringArray(questions[i].Options)

			// Insert question into the database
			if err := r.db.Create(&entities.Question{
				ID:            questions[i].ID,
				QuizID:        questions[i].QuizID,
				Text:          questions[i].Text,
				Options:       optionsArray, // Store options as Postgres array directly
				CorrectAnswer: questions[i].CorrectAnswer,
			}).Error; err != nil {
				return err
			}
		}
		// If the question already exists, you may choose to skip insertion or update it.
		// You can implement this logic based on your requirements.
	}
	return nil
}

func (r *quizRepository) GetQuiz(id string) (*entities.Quiz, error) {
	var quiz entities.Quiz
	if err := r.db.Where("id = ?", id).First(&quiz).Error; err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (r *quizRepository) GetQuizByClassID(classID string) ([]entities.Quiz, error) {
	var quiz []entities.Quiz
	// preload quiz
	if err := r.db.Preload("Class").Preload("Subject").Preload("Teacher").Preload("Questions").Where("class_id = ?", classID).Find(&quiz).Error; err != nil {
		return nil, err
	}

	return quiz, nil
}

func (r *quizRepository) GetQuestion(quizID string) ([]entities.Question, error) {
	var question []entities.Question
	if err := r.db.Where("quiz_id =?", quizID).First(&question).Error; err != nil {
		return nil, err
	}
	return question, nil
}

func (r *quizRepository) GetQuizBySubjectID(subjectID string) (*entities.Quiz, error) {
	var quiz entities.Quiz
	// preload quiz
	if err := r.db.Preload("Questions").Where("subject_id = ?", subjectID).First(&quiz).Error; err != nil {
		return nil, err
	}

	return &quiz, nil
}

func (r *quizRepository) GetQuizByTeacherID(teacherID string) ([]entities.Quiz, error) {
	var quiz []entities.Quiz
	// preload quiz
	if err := r.db.Preload("Class").Preload("Subject").Preload("Teacher").Preload("Questions").Where("teacher_id = ?", teacherID).Find(&quiz).Error; err != nil {
		return nil, err
	}

	return quiz, nil
}

func (r *quizRepository) CountQuestion(quizID string) (int64, error) {
	var count int64
	if err := r.db.Model(&entities.Question{}).
		Where("quiz_id = ?", quizID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *quizRepository) MatchAnswer(quizID string, answers []string) (int64, error) {
	var correctAnswers int64

	// Fetch all questions for the given quiz ID
	var questions []entities.Question
	if err := r.db.Model(&entities.Question{}).
		Where("quiz_id = ?", quizID).
		Find(&questions).Error; err != nil {
		return 0, err
	}

	// Create a map to store correct answers for each question ID
	correctAnswerMap := make(map[string]string)
	for _, question := range questions {
		correctAnswerMap[question.ID] = question.CorrectAnswer
	}

	// Loop through submitted answers and check against correct answers
	for _, submittedAnswer := range answers {
		for _, question := range questions {
			if submittedAnswer == correctAnswerMap[question.ID] {
				correctAnswers++
				break
			}
		}
	}

	return correctAnswers, nil
}
