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
	userRepository := repositories.NewUserRepository(db)
	tokenRepository := repositories.NewTokenRepository(db)
	roleRepository := repositories.NewRoleRepository(db)
	subjectRepository := repositories.NewSubjectRepository(db)
	teacherRepository := repositories.NewTeacherRepository(db)

	// Services
	authService := services.NewAuthService(userRepository, tokenRepository, roleRepository)
	adminService := services.NewAdminService(subjectRepository, teacherRepository)

	//Routes
	apiEndpoint := app.Group("/api")
	routes.AuthRoutes(apiEndpoint, authService)
	routes.AdminRoutes(apiEndpoint, adminService)

	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}

}
