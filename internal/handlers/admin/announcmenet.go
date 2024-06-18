package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (h *AdminHandler) CreateAnnouncement(ctx *fiber.Ctx) (err error) {
	var req dto.CreateAnnouncementRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	announcement := &entities.Announcement{
		Title:       req.Title,
		Information: req.Information,
		CreeatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	announcement, err = h.adminService.CreateAnnouncement(announcement)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create announcement",
		})
	}

	res := dto.AnnouncementResponse{
		ID:          announcement.ID,
		Title:       announcement.Title,
		Information: announcement.Information,
		CreatedAt:   announcement.CreatedAt.Format(time.DateTime),
		UpdatedAt:   announcement.UpdatedAt.Format(time.DateTime),
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Announcement created successfully",
		"data":    res,
	})

}
