package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type GlobalService interface {
	GetAnnouncements() ([]entities.Announcement, error)
	GetAnnouncementByID(id string) (*entities.Announcement, error)
}

type globalService struct {
	announcementRepo repositories.AnnouncementRepository
}

func NewGlobalService(announcementRepo repositories.AnnouncementRepository) *globalService {
	return &globalService{
		announcementRepo: announcementRepo,
	}
}

func (s *globalService) GetAnnouncements() ([]entities.Announcement, error) {
	announcements, err := s.announcementRepo.GetAnnouncements()
	if err != nil {
		return nil, services.HandleError(err, "Failed to get announcements", 500)
	}
	return announcements, nil
}

func (s *globalService) GetAnnouncementByID(id string) (*entities.Announcement, error) {
	announcement, err := s.announcementRepo.GetAnnouncementByID(id)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get announcement", 500)
	}

	return announcement, nil

}
