package app

import (
	"os"

	"github.com/Magetan-Boyz/Backend/internal/config/database"
	"github.com/Magetan-Boyz/Backend/internal/config/seeder"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/Magetan-Boyz/Backend/internal/routes"
	"github.com/Magetan-Boyz/Backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func Start() {

	app := fiber.New()

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	seed := seeder.Seed{DB: db}
	seed.SeedAll()

	// Repositories
	subjectRepository := repositories.NewSubjectRepository(db)
	teacherRepository := repositories.NewTeacherRepository(db)

	// Services
	adminService := services.NewAdminService(subjectRepository, teacherRepository)

	//Routes
	apiEndpoint := app.Group("/api")
	routes.AdminRoutes(apiEndpoint, adminService)

	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}

}
