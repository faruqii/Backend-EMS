package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type GlobalService interface {
	GetAnnouncements() ([]entities.Announcement, error)
	GetAnnouncementByID(id string) (*entities.Announcement, error)
	GetAllAgendas() ([]entities.Agenda, error)
	GetAgendaByID(id string) (*entities.Agenda, error)
}

type globalService struct {
	announcementRepo repositories.AnnouncementRepository
	agendaRepo       repositories.AgendaRepository
}

func NewGlobalService(announcementRepo repositories.AnnouncementRepository, agendaRepo repositories.AgendaRepository) *globalService {
	return &globalService{
		announcementRepo: announcementRepo,
		agendaRepo:       agendaRepo,
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

func (s *globalService) GetAllAgendas() ([]entities.Agenda, error) {
	agendas, err := s.agendaRepo.GetAllAgendas()
	if err != nil {
		return nil, services.HandleError(err, "Failed to get agendas", 500)
	}
	return agendas, nil
}

func (s *globalService) GetAgendaByID(id string) (*entities.Agenda, error) {
	agenda, err := s.agendaRepo.GetAgendaByID(id)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get agenda", 500)
	}

	return agenda, nil
}
