package app

import (
	"github.com/Magetan-Boyz/Backend/internal/services"
	adminSvc "github.com/Magetan-Boyz/Backend/internal/services/admin"
	globalSvc "github.com/Magetan-Boyz/Backend/internal/services/global"
	parentSvc "github.com/Magetan-Boyz/Backend/internal/services/parent"
	studentSvc "github.com/Magetan-Boyz/Backend/internal/services/student"
	teacherSvc "github.com/Magetan-Boyz/Backend/internal/services/teacher"
)

type Services struct {
	authService    services.AuthService
	adminService   adminSvc.AdminService
	teacherService teacherSvc.TeacherService
	studentService studentSvc.StudentService
	parentService  parentSvc.ParentService
	globalService  globalSvc.GlobalService
}

func initServices(repos *Repositories) *Services {
	return &Services{
		authService: services.NewAuthService(repos.userRepo, repos.tokenRepo, repos.roleRepo),
		adminService: adminSvc.NewAdminService(
			repos.subjectRepo, repos.teacherRepo,
			repos.userRepo, repos.roleRepo,
			repos.classRepo, repos.scheduleRepo,
			repos.studentRepo, repos.parentRepo,
			repos.announcementRepo, repos.agendaRepo),
		teacherService: teacherSvc.NewTeacherService(
			repos.teacherRepo, repos.scheduleRepo,
			repos.tokenRepo, repos.taskRepo,
			repos.classRepo, repos.subjectRepo,
			repos.quizRepo, repos.assignmentRepo,
			repos.attedanceRepo, repos.achivementRepo,
			repos.gradeRepo, repos.dispensationRepo,
			repos.literationRepo, repos.violationRepo,
			repos.studentRepo),
		studentService: studentSvc.NewStudentService(
			repos.scheduleRepo, repos.taskRepo,
			repos.studentRepo, repos.tokenRepo,
			repos.assignmentRepo, repos.quizRepo,
			repos.classRepo, repos.subjectRepo,
			repos.attedanceRepo, repos.achivementRepo,
			repos.gradeRepo, repos.dispensationRepo,
			repos.literationRepo, repos.violationRepo),
		parentService: parentSvc.NewParentService(
			repos.parentRepo, repos.scheduleRepo,
			repos.studentRepo, repos.tokenRepo,
			repos.assignmentRepo, repos.quizRepo,
			repos.classRepo, repos.subjectRepo,
			repos.attedanceRepo, repos.achivementRepo,
			repos.gradeRepo, repos.taskRepo,
			repos.violationRepo, repos.dispensationRepo),
		globalService: globalSvc.NewGlobalService(repos.announcementRepo),
	}
}
