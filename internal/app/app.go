package app

import (
	"os"

	"github.com/Magetan-Boyz/Backend/internal/config/database"
	"github.com/Magetan-Boyz/Backend/internal/config/seeder"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
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
	userRepo := repositories.NewUserRepository(db)
	tokenRepo := repositories.NewTokenRepository(db)
	roleRepo := repositories.NewRoleRepository(db)
	subjectRepo := repositories.NewSubjectRepository(db)
	teacherRepo := repositories.NewTeacherRepository(db)
	classRepo := repositories.NewClassRepository(db)

	// Services
	authService := services.NewAuthService(userRepo, tokenRepo, roleRepo)
	adminService := services.NewAdminService(subjectRepo, teacherRepo, userRepo, roleRepo, classRepo)

	// Middleware
	middleware := middleware.NewMiddleware(tokenRepo, roleRepo)
	
	//Routes
	api := app.Group("/api")
	routes.AuthRoutes(api, authService)
	routes.AdminRoutes(api, adminService, middleware)

	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}

}
