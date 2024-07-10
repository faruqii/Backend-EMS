package app

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
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
	announcementRepo repositories.AnnouncementRepository
	agendaRepo       repositories.AgendaRepository
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
		announcementRepo: repositories.NewAnnouncementRepository(db),
		agendaRepo:       repositories.NewAgendaRepository(db),
	}
}
