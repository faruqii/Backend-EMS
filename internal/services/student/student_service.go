package service

import "github.com/Magetan-Boyz/Backend/internal/domain/repositories"

type StudentService interface {
	StudentTaskService
	StudentAssignmentService
	StudentQuizService
}

type studentService struct {
	scheduleRepo   repositories.ScheduleRepository
	taskRepo       repositories.TaskRepository
	studentRepo    repositories.StudentRepository
	tokenRepo      repositories.TokenRepository
	assignmentRepo repositories.AssignmentRepository
	quizRepo       repositories.QuizRepository
}

func NewStudentService(
	scheduleRepo repositories.ScheduleRepository,
	taskRepo repositories.TaskRepository,
	studentRepo repositories.StudentRepository,
	tokenRepo repositories.TokenRepository,
	assignmentRepo repositories.AssignmentRepository,
	quizRepo repositories.QuizRepository,
) *studentService {
	return &studentService{
		scheduleRepo:   scheduleRepo,
		taskRepo:       taskRepo,
		studentRepo:    studentRepo,
		tokenRepo:      tokenRepo,
		assignmentRepo: assignmentRepo,
		quizRepo:       quizRepo,
	}
}
