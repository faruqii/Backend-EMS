package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

var location *time.Location

func init() {
	var err error
	location, err = time.LoadLocation("Asia/Jakarta") // Set to your desired timezone
	if err != nil {
		panic(err)
	}
}

func formatTimeWithLocation(t time.Time) string {
	return t.In(location).Format("2006-01-02 15:04:05")
}

func (h *GlobalHandler) GetAnnouncements(ctx *fiber.Ctx) (err error) {
	announcements, err := h.globalService.GetAnnouncements()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get announcements",
		})
	}

	var res []dto.AnnouncementResponse
	for _, announcement := range announcements {
		res = append(res, dto.AnnouncementResponse{
			ID:          announcement.ID,
			Title:       announcement.Title,
			Information: announcement.Information,
			CreatedAt:   formatTimeWithLocation(announcement.CreatedAt),
			UpdatedAt:   formatTimeWithLocation(announcement.UpdatedAt),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Announcements fetched successfully",
		"data":    res,
	})
}

func (h *GlobalHandler) GetAnnouncementByID(ctx *fiber.Ctx) (err error) {
	announcementID := ctx.Params("id")
	announcement, err := h.globalService.GetAnnouncementByID(announcementID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get announcement",
		})
	}

	res := dto.AnnouncementResponse{
		ID:          announcement.ID,
		Title:       announcement.Title,
		Information: announcement.Information,
		CreatedAt:   formatTimeWithLocation(announcement.CreatedAt),
		UpdatedAt:   formatTimeWithLocation(announcement.UpdatedAt),
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Announcement fetched successfully",
		"data":    res,
	})
}

func (h *GlobalHandler) GetAllAgendas(ctx *fiber.Ctx) (err error) {
	agendas, err := h.globalService.GetAllAgendas()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get agendas",
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
		"message": "Agendas fetched successfully",
		"data":    response,
	})
}

func (h *GlobalHandler) GetAgendaByID(ctx *fiber.Ctx) (err error) {
	agendaID := ctx.Params("agendaID")

	agenda, err := h.globalService.GetAgendaByID(agendaID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get agenda",
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
		"message": "Agenda fetched successfully",
		"data":    response,
	})
}