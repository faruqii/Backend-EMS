package app

import (
	"log"
	"os"

	"github.com/Magetan-Boyz/Backend/internal/config/database"
	"github.com/Magetan-Boyz/Backend/internal/config/seeder"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	"github.com/Magetan-Boyz/Backend/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App, services *Services, mw *middleware.Middleware) {
	api := app.Group("/api")
	routes.AuthRoutes(api, services.authService, mw)
	routes.AdminRoutes(api, services.adminService, mw)
	routes.TeacherRoutes(api, services.teacherService, mw)
	routes.StudentRoutes(api, services.studentService, mw)
	routes.ParentRoutes(api, services.parentService, mw)
	routes.GlobalRoutes(api, services.globalService)
}

func Start() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS, PATCH, HEAD",
		AllowCredentials: true,
	}))

	// Database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Seed data
	seed := seeder.Seed{DB: db}
	seed.SeedAll()

	// Repositories
	repos := initRepositories(db)

	// Services
	services := initServices(repos)

	// Middleware
	mw := middleware.NewMiddleware(repos.tokenRepo, repos.roleRepo)

	// Routes
	setupRoutes(app, services, mw)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port if not specified
	}

	err = app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
