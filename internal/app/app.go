package app

import (
	"os"

	"github.com/Magetan-Boyz/Backend/internal/config/database"
	"github.com/Magetan-Boyz/Backend/internal/config/seeder"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	"github.com/Magetan-Boyz/Backend/internal/routes"
	"github.com/Magetan-Boyz/Backend/internal/services"
	adminSvc "github.com/Magetan-Boyz/Backend/internal/services/admin"
	studentSvc "github.com/Magetan-Boyz/Backend/internal/services/student"
	teacherSvc "github.com/Magetan-Boyz/Backend/internal/services/teacher"
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
	scheduleRepo := repositories.NewScheduleRepository(db)
	studentRepo := repositories.NewStudentRepository(db)
	taskRepo := repositories.NewTaskRepository(db)
	assisgnmentRepo := repositories.NewAssignmentRepository(db)
	quizRepo := repositories.NewQuizRepository(db)

	// Services
	authService := services.NewAuthService(userRepo, tokenRepo, roleRepo)
	adminService := adminSvc.NewAdminService(subjectRepo, teacherRepo, userRepo, roleRepo, classRepo, scheduleRepo, studentRepo)
	teacherService := teacherSvc.NewTeacherService(teacherRepo, scheduleRepo, tokenRepo, taskRepo, classRepo, subjectRepo, quizRepo)
	studentService := studentSvc.NewStudentService(scheduleRepo, taskRepo, studentRepo, tokenRepo, assisgnmentRepo)

	// Middleware
	middleware := middleware.NewMiddleware(tokenRepo, roleRepo)

	//Routes
	api := app.Group("/api")
	routes.AuthRoutes(api, authService, middleware)
	routes.AdminRoutes(api, adminService, middleware)
	routes.TeacherRoutes(api, teacherService, middleware)
	routes.StudentRoutes(api, studentService, middleware)

	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}

}
