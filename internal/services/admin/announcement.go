package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type AdminAnnouncementService interface {
	CreateAnnouncement(announcement *entities.Announcement) (*entities.Announcement, error)
	GetAnnouncements() ([]entities.Announcement, error)
	GetAnnouncementByID(announcementID string) (*entities.Announcement, error)
	UpdateAnnouncement(announcement *entities.Announcement) (*entities.Announcement, error)
	DeleteAnnouncement(announcementID string) error
}

func (s *adminService) CreateAnnouncement(announcement *entities.Announcement) (*entities.Announcement, error) {
	announcement, err := s.announcementRepo.CreateAnnouncement(announcement)
	if err != nil {
		return nil, services.HandleError(err, "Failed to create announcement", 500)
	}

	return announcement, nil
}

func (s *adminService) GetAnnouncements() ([]entities.Announcement, error) {
	announcements, err := s.announcementRepo.GetAnnouncements()
	if err != nil {
		return nil, services.HandleError(err, "Failed to get announcements", 500)
	}

	return announcements, nil
}

func (s *adminService) GetAnnouncementByID(announcementID string) (*entities.Announcement, error) {
	announcement, err := s.announcementRepo.GetAnnouncementByID(announcementID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get announcement", 500)
	}

	return announcement, nil
}

func (s *adminService) UpdateAnnouncement(announcement *entities.Announcement) (*entities.Announcement, error) {
	announcement, err := s.announcementRepo.UpdateAnnouncement(announcement)
	if err != nil {
		return nil, services.HandleError(err, "Failed to update announcement", 500)
	}

	return announcement, nil
}

func (s *adminService) DeleteAnnouncement(announcementID string) error {
	err := s.announcementRepo.DeleteAnnouncement(announcementID)
	if err != nil {
		return services.HandleError(err, "Failed to delete announcement", 500)
	}

	return nil
}
