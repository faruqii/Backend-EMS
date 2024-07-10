package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
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

func (h *AdminHandler) CreateAnnouncement(ctx *fiber.Ctx) (err error) {
	var req dto.CreateAnnouncementRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	now := time.Now().In(location)
	announcement := &entities.Announcement{
		Title:       req.Title,
		Information: req.Information,
		CreatedAt:   now,
		UpdatedAt:   now,
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
		CreatedAt:   formatTimeWithLocation(announcement.CreatedAt),
		UpdatedAt:   formatTimeWithLocation(announcement.UpdatedAt),
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Announcement created successfully",
		"data":    res,
	})
}

func (h *AdminHandler) GetAnnouncements(ctx *fiber.Ctx) (err error) {
	announcements, err := h.adminService.GetAnnouncements()
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
		"data": res,
	})
}

func (h *AdminHandler) GetAnnouncementByID(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")
	announcement, err := h.adminService.GetAnnouncementByID(id)
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
		"data": res,
	})
}

func (h *AdminHandler) UpdateAnnouncement(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")
	var req dto.UpdateAnnouncementRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	announcement, err := h.adminService.GetAnnouncementByID(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get announcement",
		})
	}

	announcement.Title = req.Title
	announcement.Information = req.Information
	announcement.UpdatedAt = time.Now().In(location)

	announcement, err = h.adminService.UpdateAnnouncement(announcement)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update announcement",
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
		"message": "Announcement updated successfully",
		"data":    res,
	})
}

func (h *AdminHandler) DeleteAnnouncement(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	err = h.adminService.DeleteAnnouncement(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete announcement",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Announcement deleted successfully",
	})
}
