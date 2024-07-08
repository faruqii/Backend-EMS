package repositories

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"gorm.io/gorm"
)

type AgendaRepository interface {
	CreateAgenda(agenda *entities.Agenda) error
	GetAgendaByID(id string) (*entities.Agenda, error)
	GetAllAgendas() ([]entities.Agenda, error)
	UpdateAgenda(id string, updatedAgenda *entities.Agenda) error
	DeleteAgenda(id string) error
}

type agendaRepository struct {
	db *gorm.DB
}

func NewAgendaRepository(db *gorm.DB) *agendaRepository {
	return &agendaRepository{db: db}
}

func (r *agendaRepository) CreateAgenda(agenda *entities.Agenda) error {
	if err := r.db.Create(agenda).Error; err != nil {
		return err
	}
	return nil
}

func (r *agendaRepository) GetAgendaByID(id string) (*entities.Agenda, error) {
	var agenda entities.Agenda
	if err := r.db.Where("id = ?", id).First(&agenda).Error; err != nil {
		return nil, err
	}
	return &agenda, nil
}

func (r *agendaRepository) GetAllAgendas() ([]entities.Agenda, error) {
	var agendas []entities.Agenda
	if err := r.db.Find(&agendas).Error; err != nil {
		return nil, err
	}
	return agendas, nil
}

func (r *agendaRepository) UpdateAgenda(id string, updatedAgenda *entities.Agenda) error {
	if err := r.db.Model(&entities.Agenda{}).Where("id = ?", id).Updates(updatedAgenda).Error; err != nil {
		return err
	}
	return nil
}

func (r *agendaRepository) DeleteAgenda(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.Agenda{}).Error; err != nil {
		return err
	}
	return nil
}
