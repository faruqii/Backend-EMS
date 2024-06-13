package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type AchivementRepository interface {
	InsertAchivement(achivement *entities.StudentAchivement) (*entities.StudentAchivement, error)
	GetAchivementByID(achivementID string) (*entities.StudentAchivement, error)
	GetAllAchivement(studentID string) ([]entities.StudentAchivement, error)
	UpdateAchivement(achivement *entities.StudentAchivement) (*entities.StudentAchivement, error)
	DeleteAchivement(achivementID string) error
}

type achivementRepository struct {
	db *gorm.DB
}

func NewAchivementRepository(db *gorm.DB) AchivementRepository {
	return &achivementRepository{db: db}
}

func (r *achivementRepository) InsertAchivement(achivement *entities.StudentAchivement) (*entities.StudentAchivement, error) {
	if err := r.db.Create(achivement).Error; err != nil {
		return nil, err
	}

	return achivement, nil
}

func (r *achivementRepository) GetAchivementByID(achivementID string) (*entities.StudentAchivement, error) {
	achivement := entities.StudentAchivement{}
	if err := r.db.Preload("Student").Where("id = ?", achivementID).First(&achivement).Error; err != nil {
		return nil, err
	}

	return &achivement, nil
}

func (r *achivementRepository) GetAllAchivement(studentID string) ([]entities.StudentAchivement, error) {
	achivements := []entities.StudentAchivement{}
	if err := r.db.Preload("Student").Where("student_id = ?", studentID).Find(&achivements).Error; err != nil {
		return nil, err
	}

	return achivements, nil
}

func (r *achivementRepository) UpdateAchivement(achivement *entities.StudentAchivement) (*entities.StudentAchivement, error) {
	if err := r.db.Save(achivement).Error; err != nil {
		return nil, err
	}

	return achivement, nil
}

func (r *achivementRepository) DeleteAchivement(achivementID string) error {
	if err := r.db.Where("id = ?", achivementID).Delete(&entities.StudentAchivement{}).Error; err != nil {
		return err
	}

	return nil
}
