package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type AgendaService interface {
	CreateAgenda(agenda *entities.Agenda) error
	GetAgendaByID(id string) (*entities.Agenda, error)
	GetAllAgendas() ([]entities.Agenda, error)
	UpdateAgenda(id string, updatedAgenda *entities.Agenda) error
	DeleteAgenda(id string) error
}

func (s *adminService) CreateAgenda(agenda *entities.Agenda) error {
	err := s.agendaRepo.CreateAgenda(agenda)
	if err != nil {
		return services.HandleError(err, "Failed to create agenda", 500)
	}

	return nil
}

func (s *adminService) GetAgendaByID(id string) (*entities.Agenda, error) {
	agenda, err := s.agendaRepo.GetAgendaByID(id)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch agenda", 500)
	}

	return agenda, nil
}

func (s *adminService) GetAllAgendas() ([]entities.Agenda, error) {
	agendas, err := s.agendaRepo.GetAllAgendas()
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch agendas", 500)
	}

	return agendas, nil
}

func (s *adminService) UpdateAgenda(id string, updatedAgenda *entities.Agenda) error {
	err := s.agendaRepo.UpdateAgenda(id, updatedAgenda)
	if err != nil {
		return services.HandleError(err, "Failed to update agenda", 500)
	}

	return nil
}

func (s *adminService) DeleteAgenda(id string) error {
	err := s.agendaRepo.DeleteAgenda(id)
	if err != nil {
		return services.HandleError(err, "Failed to delete agenda", 500)
	}

	return nil
}
