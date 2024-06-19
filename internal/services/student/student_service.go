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
	StudentGradeService
	StudentDispensationService
	StudentLiterationService
	StudentProfile
	StudentViolationService
}

type studentService struct {
	scheduleRepo     repositories.ScheduleRepository
	taskRepo         repositories.TaskRepository
	studentRepo      repositories.StudentRepository
	tokenRepo        repositories.TokenRepository
	assignmentRepo   repositories.AssignmentRepository
	quizRepo         repositories.QuizRepository
	classRepo        repositories.ClassRepository
	subjectRepo      repositories.SubjectRepository
	attedanceRepo    repositories.AttedanceRepository
	achivementRepo   repositories.AchivementRepository
	gradeRepo        repositories.GradeRepository
	dispensationRepo repositories.DispensationRepository
	literationRepo   repositories.LiterationRepository
	violationRepo    repositories.ViolationRepository
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
	gradeRepo repositories.GradeRepository,
	dispensationRepo repositories.DispensationRepository,
	literationRepo repositories.LiterationRepository,
	violationRepo repositories.ViolationRepository,
) *studentService {
	return &studentService{
		scheduleRepo:     scheduleRepo,
		taskRepo:         taskRepo,
		studentRepo:      studentRepo,
		tokenRepo:        tokenRepo,
		assignmentRepo:   assignmentRepo,
		quizRepo:         quizRepo,
		classRepo:        classRepo,
		subjectRepo:      subjectRepo,
		attedanceRepo:    attedanceRepo,
		achivementRepo:   achivementRepo,
		gradeRepo:        gradeRepo,
		dispensationRepo: dispensationRepo,
		literationRepo:   literationRepo,
		violationRepo:    violationRepo,
	}
}
