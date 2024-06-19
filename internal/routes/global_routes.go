package routes

import (
	handlers "github.com/Magetan-Boyz/Backend/internal/handlers/global"
	service "github.com/Magetan-Boyz/Backend/internal/services/global"
	"github.com/gofiber/fiber/v2"
)

func GlobalRoutes(router fiber.Router, globalService service.GlobalService) {
	globalHandler := handlers.NewGlobalHandler(globalService)

	globalRoutes := router.Group("/global")

	globalRoutes.Get("/announcements", globalHandler.GetAnnouncements)
	globalRoutes.Get("/announcements/:id", globalHandler.GetAnnouncementByID)
}
