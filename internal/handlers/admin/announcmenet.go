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
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
			CreatedAt:   announcement.CreatedAt.Format(time.DateTime),
			UpdatedAt:   announcement.UpdatedAt.Format(time.DateTime),
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
		CreatedAt:   announcement.CreatedAt.Format(time.DateTime),
		UpdatedAt:   announcement.UpdatedAt.Format(time.DateTime),
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
	announcement.UpdatedAt = time.Now()

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
		CreatedAt:   announcement.CreatedAt.Format(time.DateTime),
		UpdatedAt:   announcement.UpdatedAt.Format(time.DateTime),
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
