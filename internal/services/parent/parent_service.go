package service

import "github.com/Magetan-Boyz/Backend/internal/domain/repositories"

type ParentService interface {
	ParentStudentSchedule
	ParentTaskService
	ParentStudentAchievementService
	ParentStudentQuizService
	ParentStudentGradeService
	ParentTokenService
	ParentStudentViolationService
	ParentStudentService
}

type parentService struct {
	parentRepo     repositories.ParentRepository
	scheduleRepo   repositories.ScheduleRepository
	studentRepo    repositories.StudentRepository
	tokenRepo      repositories.TokenRepository
	assignmentRepo repositories.AssignmentRepository
	quizRepo       repositories.QuizRepository
	classRepo      repositories.ClassRepository
	subjectRepo    repositories.SubjectRepository
	attedanceRepo  repositories.AttedanceRepository
	achivementRepo repositories.AchivementRepository
	gradeRepo      repositories.GradeRepository
	taskRepo       repositories.TaskRepository
	violationRepo  repositories.ViolationRepository
}

func NewParentService(
	parentRepo repositories.ParentRepository,
	scheduleRepo repositories.ScheduleRepository,
	studentRepo repositories.StudentRepository,
	tokenRepo repositories.TokenRepository,
	assignmentRepo repositories.AssignmentRepository,
	quizRepo repositories.QuizRepository,
	classRepo repositories.ClassRepository,
	subjectRepo repositories.SubjectRepository,
	attedanceRepo repositories.AttedanceRepository,
	achivementRepo repositories.AchivementRepository,
	gradeRepo repositories.GradeRepository,
	taskRepo repositories.TaskRepository,
	violationRepo repositories.ViolationRepository,
) *parentService {
	return &parentService{
		parentRepo:     parentRepo,
		scheduleRepo:   scheduleRepo,
		studentRepo:    studentRepo,
		tokenRepo:      tokenRepo,
		assignmentRepo: assignmentRepo,
		quizRepo:       quizRepo,
		classRepo:      classRepo,
		subjectRepo:    subjectRepo,
		attedanceRepo:  attedanceRepo,
		achivementRepo: achivementRepo,
		gradeRepo:      gradeRepo,
		taskRepo:       taskRepo,
		violationRepo:  violationRepo,
	}
}
