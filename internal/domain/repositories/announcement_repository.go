package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type AnnouncementRepository interface {
	CreateAnnouncement(announcement *entities.Announcement) (*entities.Announcement, error)
	GetAnnouncements() ([]entities.Announcement, error)
	GetAnnouncementByID(announcementID string) (*entities.Announcement, error)
	UpdateAnnouncement(announcement *entities.Announcement) (*entities.Announcement, error)
	DeleteAnnouncement(announcementID string) error
}

type announcementRepository struct {
	db *gorm.DB
}

func NewAnnouncementRepository(db *gorm.DB) AnnouncementRepository {
	return &announcementRepository{
		db: db,
	}
}

func (r *announcementRepository) CreateAnnouncement(announcement *entities.Announcement) (*entities.Announcement, error) {
	if err := r.db.Create(announcement).Error; err != nil {
		return nil, err
	}

	return announcement, nil
}

func (r *announcementRepository) GetAnnouncements() ([]entities.Announcement, error) {
	var announcements []entities.Announcement

	if err := r.db.Find(&announcements).Error; err != nil {
		return nil, err
	}

	return announcements, nil
}

func (r *announcementRepository) GetAnnouncementByID(announcementID string) (*entities.Announcement, error) {
	var announcement entities.Announcement

	if err := r.db.Where("id = ?", announcementID).First(&announcement).Error; err != nil {
		return nil, err
	}

	return &announcement, nil
}

func (r *announcementRepository) UpdateAnnouncement(announcement *entities.Announcement) (*entities.Announcement, error) {
	if err := r.db.Where("id = ?", announcement.ID).Updates(announcement).Error; err != nil {
		return nil, err
	}

	return announcement, nil

}

func (r *announcementRepository) DeleteAnnouncement(announcementID string) error {
	if err := r.db.Where("id = ?", announcementID).Delete(&entities.Announcement{}).Error; err != nil {
		return err
	}

	return nil
}
