package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

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
			CreatedAt:   announcement.CreatedAt.Format(time.DateTime),
			UpdatedAt:   announcement.UpdatedAt.Format(time.DateTime),
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
		CreatedAt:   announcement.CreatedAt.Format(time.DateTime),
		UpdatedAt:   announcement.UpdatedAt.Format(time.DateTime),
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Announcement fetched successfully",
		"data":    res,
	})
}
