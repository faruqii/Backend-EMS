package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/repositories"
)

type TeacherService interface {
	TeacherScheduleService
	TeacherTaskService
	QuizService
}

type teacherService struct {
	teacherRepo  repositories.TeacherRepository
	scheduleRepo repositories.ScheduleRepository
	tokenRepo    repositories.TokenRepository
	taskRepo     repositories.TaskRepository
	classRepo    repositories.ClassRepository
	subjectRepo  repositories.SubjectRepository
	quizRepo     repositories.QuizRepository
}

func NewTeacherService(
	teacherRepo repositories.TeacherRepository,
	scheduleRepo repositories.ScheduleRepository,
	tokenRepo repositories.TokenRepository,
	taskRepo repositories.TaskRepository,
	classRepo repositories.ClassRepository,
	subjectRepo repositories.SubjectRepository,
	quizRepo repositories.QuizRepository,
) *teacherService {
	return &teacherService{
		teacherRepo:  teacherRepo,
		scheduleRepo: scheduleRepo,
		tokenRepo:    tokenRepo,
		taskRepo:     taskRepo,
		classRepo:    classRepo,
		subjectRepo:  subjectRepo,
		quizRepo:     quizRepo,
	}
}
