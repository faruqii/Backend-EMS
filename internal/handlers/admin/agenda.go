package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (h *AdminHandler) CreateAgenda(ctx *fiber.Ctx) error {
	var req dto.CreateAgendaRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// load location
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// parse date
	date, err := time.ParseInLocation(time.DateOnly, req.Date, loc)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// parse start time and end time
	startTime, err := time.ParseInLocation(time.TimeOnly, req.StartTime, loc)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	endTime, err := time.ParseInLocation(time.TimeOnly, req.EndTime, loc)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	agenda := entities.Agenda{
		Title:        req.Title,
		Date:         date,
		StartTime:    startTime,
		EndTime:      endTime,
		TypeOfAgenda: req.TypeOfAgenda,
		Location:     req.Location,
		Description:  req.Description,
	}

	err = h.adminService.CreateAgenda(&agenda)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Agenda created successfully",
	})
}

func (h *AdminHandler) UpdateAgenda(ctx *fiber.Ctx) error {
	agendaID := ctx.Params("agendaID")

	if agendaID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "agendaID is required",
		})
	}

	var req dto.UpdateAgendaRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	agenda, err := h.adminService.GetAgendaByID(agendaID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// load location
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// parse date
	date, err := time.ParseInLocation(time.DateOnly, req.Date, loc)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// parse start time and end time
	startTime, err := time.ParseInLocation(time.TimeOnly, req.StartTime, loc)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	endTime, err := time.ParseInLocation(time.TimeOnly, req.EndTime, loc)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	agenda.Title = req.Title
	agenda.Date = date
	agenda.StartTime = startTime
	agenda.EndTime = endTime
	agenda.TypeOfAgenda = req.TypeOfAgenda
	agenda.Location = req.Location
	agenda.Description = req.Description

	err = h.adminService.UpdateAgenda(agenda.ID, agenda)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Agenda updated successfully",
	})

}

func (h *AdminHandler) DeleteAgenda(ctx *fiber.Ctx) error {
	agendaID := ctx.Params("agendaID")

	if agendaID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "agendaID is required",
		})
	}

	err := h.adminService.DeleteAgenda(agendaID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Agenda deleted successfully",
	})
}

func (h *AdminHandler) GetAgendaByID(ctx *fiber.Ctx) error {
	agendaID := ctx.Params("agendaID")

	if agendaID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "agendaID is required",
		})
	}

	agenda, err := h.adminService.GetAgendaByID(agendaID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.AgendaResponse{
		ID:           agenda.ID,
		Title:        agenda.Title,
		Date:         agenda.Date.Format(time.DateOnly),
		StartTime:    agenda.StartTime.Format(time.TimeOnly),
		EndTime:      agenda.EndTime.Format(time.TimeOnly),
		TypeOfAgenda: agenda.TypeOfAgenda,
		Location:     agenda.Location,
		Description:  agenda.Description,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Agenda retrieved successfully",
		"data":    response,
	})

}

func (h *AdminHandler) GetAllAgenda(ctx *fiber.Ctx) error {
	agendas, err := h.adminService.GetAllAgendas()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []dto.AgendaResponse

	for _, agenda := range agendas {
		response = append(response, dto.AgendaResponse{
			ID:           agenda.ID,
			Title:        agenda.Title,
			Date:         agenda.Date.Format(time.DateOnly),
			StartTime:    agenda.StartTime.Format(time.TimeOnly),
			EndTime:      agenda.EndTime.Format(time.TimeOnly),
			TypeOfAgenda: agenda.TypeOfAgenda,
			Location:     agenda.Location,
			Description:  agenda.Description,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Agendas retrieved successfully",
		"data":    response,
	})
}
