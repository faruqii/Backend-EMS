package service

import "github.com/Magetan-Boyz/Backend/internal/domain/repositories"

type StudentService interface {
	StudentTaskService
	StudentAssignmentService
	StudentQuizService
	StudentClassService
	StudentAttendanceService
	StudentSubjectRepository
	StudentAchivementService
}

type studentService struct {
	scheduleRepo   repositories.ScheduleRepository
	taskRepo       repositories.TaskRepository
	studentRepo    repositories.StudentRepository
	tokenRepo      repositories.TokenRepository
	assignmentRepo repositories.AssignmentRepository
	quizRepo       repositories.QuizRepository
	classRepo      repositories.ClassRepository
	subjectRepo    repositories.SubjectRepository
	attedanceRepo  repositories.AttedanceRepository
	achivementRepo repositories.AchivementRepository
}

func NewStudentService(
	scheduleRepo repositories.ScheduleRepository,
	taskRepo repositories.TaskRepository,
	studentRepo repositories.StudentRepository,
	tokenRepo repositories.TokenRepository,
	assignmentRepo repositories.AssignmentRepository,
	quizRepo repositories.QuizRepository,
	classRepo repositories.ClassRepository,
	subjectRepo repositories.SubjectRepository,
	attedanceRepo repositories.AttedanceRepository,
	achivementRepo repositories.AchivementRepository,
) *studentService {
	return &studentService{
		scheduleRepo:   scheduleRepo,
		taskRepo:       taskRepo,
		studentRepo:    studentRepo,
		tokenRepo:      tokenRepo,
		assignmentRepo: assignmentRepo,
		quizRepo:       quizRepo,
		classRepo:      classRepo,
		subjectRepo:    subjectRepo,
		attedanceRepo:  attedanceRepo,
		achivementRepo: achivementRepo,
	}
}
