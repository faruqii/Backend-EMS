package app

import (
	"log"
	"os"

	"github.com/Magetan-Boyz/Backend/internal/config/database"
	"github.com/Magetan-Boyz/Backend/internal/config/seeder"
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	"github.com/Magetan-Boyz/Backend/internal/routes"
	"github.com/Magetan-Boyz/Backend/internal/services"
	adminSvc "github.com/Magetan-Boyz/Backend/internal/services/admin"
	parentSvc "github.com/Magetan-Boyz/Backend/internal/services/parent"
	studentSvc "github.com/Magetan-Boyz/Backend/internal/services/student"
	teacherSvc "github.com/Magetan-Boyz/Backend/internal/services/teacher"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repositories struct {
	userRepo         repositories.UserRepository
	tokenRepo        repositories.TokenRepository
	roleRepo         repositories.RoleRepository
	subjectRepo      repositories.SubjectRepository
	teacherRepo      repositories.TeacherRepository
	classRepo        repositories.ClassRepository
	scheduleRepo     repositories.ScheduleRepository
	studentRepo      repositories.StudentRepository
	taskRepo         repositories.TaskRepository
	assignmentRepo   repositories.AssignmentRepository
	quizRepo         repositories.QuizRepository
	attedanceRepo    repositories.AttedanceRepository
	parentRepo       repositories.ParentRepository
	achivementRepo   repositories.AchivementRepository
	gradeRepo        repositories.GradeRepository
	dispensationRepo repositories.DispensationRepository
	literationRepo   repositories.LiterationRepository
	violationRepo    repositories.ViolationRepository
}

func initRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		userRepo:         repositories.NewUserRepository(db),
		tokenRepo:        repositories.NewTokenRepository(db),
		roleRepo:         repositories.NewRoleRepository(db),
		subjectRepo:      repositories.NewSubjectRepository(db),
		teacherRepo:      repositories.NewTeacherRepository(db),
		classRepo:        repositories.NewClassRepository(db),
		scheduleRepo:     repositories.NewScheduleRepository(db),
		studentRepo:      repositories.NewStudentRepository(db),
		taskRepo:         repositories.NewTaskRepository(db),
		assignmentRepo:   repositories.NewAssignmentRepository(db),
		quizRepo:         repositories.NewQuizRepository(db),
		attedanceRepo:    repositories.NewAttedanceRepository(db),
		parentRepo:       repositories.NewParentRepository(db),
		achivementRepo:   repositories.NewAchivementRepository(db),
		gradeRepo:        repositories.NewGradeRepository(db),
		dispensationRepo: repositories.NewDispensationRepository(db),
		literationRepo:   repositories.NewLiterationRepository(db),
		violationRepo:    repositories.NewViolationRepository(db),
	}
}

type Services struct {
	authService    services.AuthService
	adminService   adminSvc.AdminService
	teacherService teacherSvc.TeacherService
	studentService studentSvc.StudentService
	parentService  parentSvc.ParentService
}

func initServices(repos *Repositories) *Services {
	return &Services{
		authService: services.NewAuthService(repos.userRepo, repos.tokenRepo, repos.roleRepo),
		adminService: adminSvc.NewAdminService(
			repos.subjectRepo, repos.teacherRepo,
			repos.userRepo, repos.roleRepo,
			repos.classRepo, repos.scheduleRepo,
			repos.studentRepo, repos.parentRepo),
		teacherService: teacherSvc.NewTeacherService(
			repos.teacherRepo, repos.scheduleRepo,
			repos.tokenRepo, repos.taskRepo,
			repos.classRepo, repos.subjectRepo,
			repos.quizRepo, repos.assignmentRepo,
			repos.attedanceRepo, repos.achivementRepo,
			repos.gradeRepo, repos.dispensationRepo,
			repos.literationRepo, repos.violationRepo),
		studentService: studentSvc.NewStudentService(
			repos.scheduleRepo, repos.taskRepo,
			repos.studentRepo, repos.tokenRepo,
			repos.assignmentRepo, repos.quizRepo,
			repos.classRepo, repos.subjectRepo,
			repos.attedanceRepo, repos.achivementRepo,
			repos.gradeRepo, repos.dispensationRepo, repos.literationRepo),
		parentService: parentSvc.NewParentService(
			repos.parentRepo, repos.scheduleRepo,
			repos.studentRepo, repos.tokenRepo,
			repos.assignmentRepo, repos.quizRepo,
			repos.classRepo, repos.subjectRepo,
			repos.attedanceRepo, repos.achivementRepo,
			repos.gradeRepo, repos.taskRepo),
	}
}

func setupRoutes(app *fiber.App, services *Services, mw *middleware.Middleware) {
	api := app.Group("/api")
	routes.AuthRoutes(api, services.authService, mw)
	routes.AdminRoutes(api, services.adminService, mw)
	routes.TeacherRoutes(api, services.teacherService, mw)
	routes.StudentRoutes(api, services.studentService, mw)
	routes.ParentRoutes(api, services.parentService, mw)
}

func Start() {
	app := fiber.New()

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
