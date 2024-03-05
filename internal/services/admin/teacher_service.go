package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"golang.org/x/crypto/bcrypt"
)

type AdminTeacherService interface {
	CreateTeacher(teacher *entities.Teacher) error
	GetAllTeacher() ([]entities.Teacher, error)
	AssignTeacherToSubject(teacherID, SubjectID string) error
	FindTeacherByID(id string) (*entities.Teacher, error)
	GetTeachersBySubjectID(subjectID string) ([]dto.TeacherSubjectResponse, error)
	GetTeacherSubjects(teacherID string) ([]dto.TeacherSubjectsResponse, error)
	UpdateTeacherHomeroomStatus(teacherID string, isHomeroom bool) error
}

func (s *adminService) CreateTeacher(teacher *entities.Teacher) error {
	_, err := s.userRepo.FindByUsername(teacher.User.Username)
	if err == nil {
		return s.handleError(err, "Username already exist", 400)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(teacher.User.Password), bcrypt.MinCost)
	if err != nil {
		return s.handleError(err, "Failed to hash password", 500)
	}

	teacher.User.Password = string(hashedPassword)

	err = s.teacherRepo.Create(teacher)
	if err != nil {
		return s.handleError(err, "Failed to create teacher", 500)
	}

	err = s.roleRepo.AssignUserRole(teacher.User.ID, "teacher")
	return s.handleError(err, "Failed to assign role to teacher", 500)
}

func (s *adminService) GetAllTeacher() ([]entities.Teacher, error) {
	teachers, err := s.teacherRepo.GetAll()
	return teachers, s.handleError(err, "Failed to fetch teachers", 500)
}

func (s *adminService) AssignTeacherToSubject(teacherID, SubjectID string) error {
	isAssigned, err := s.subjectRepo.IsTeacherAssignedToSubject(teacherID, SubjectID)
	if err != nil {
		return s.handleError(err, "Failed to check if teacher is assigned to subject", 500)
	}
	if isAssigned {
		return s.handleError(err, "Teacher already assigned to subject", 400)
	}

	err = s.subjectRepo.AssignTeacherToSubject(teacherID, SubjectID)
	return s.handleError(err, "Failed to assign teacher to subject", 500)
}

func (s *adminService) FindTeacherByID(id string) (*entities.Teacher, error) {
	teacher, err := s.teacherRepo.FindByID(id)
	return teacher, s.handleError(err, "Failed to fetch teacher", 500)
}

func (s *adminService) GetTeachersBySubjectID(subjectID string) ([]dto.TeacherSubjectResponse, error) {
	teacherSubjects, err := s.subjectRepo.GetTeachersBySubjectID(subjectID)
	if err != nil {
		return nil, s.handleError(err, "Failed to fetch teachers", 500)
	}

	var teachers []dto.TeacherSubjectResponse
	for _, ts := range teacherSubjects {
		teachers = append(teachers, dto.TeacherSubjectResponse{
			TeacherName: ts.Teacher.Name,
			SubjectName: ts.Subject.Name,
		})
	}
	return teachers, nil
}

func (s *adminService) GetTeacherSubjects(teacherID string) ([]dto.TeacherSubjectsResponse, error) {
	teacherSubjects, err := s.subjectRepo.GetTeacherSubjects(teacherID)
	if err != nil {
		return nil, s.handleError(err, "Failed to fetch teacher subjects", 500)
	}

	var subjects []dto.TeacherSubjectsResponse

	if len(teacherSubjects) == 0 {
		subjects = append(subjects, dto.TeacherSubjectsResponse{
			TeacherName: "",
			SubjectName: []string{},
		})
	} else {
		subjectMap := make(map[string]bool)
		for _, ts := range teacherSubjects {
			subjectMap[ts.Subject.Name] = true
		}
		firstTeacherName := teacherSubjects[0].Teacher.Name
		subjects = append(subjects, dto.TeacherSubjectsResponse{
			TeacherName: firstTeacherName,
			SubjectName: make([]string, 0, len(subjectMap)),
		})
		for subject := range subjectMap {
			subjects[0].SubjectName = append(subjects[0].SubjectName, subject)
		}
	}

	return subjects, nil
}

func (s *adminService) UpdateTeacherHomeroomStatus(teacherID string, isHomeroom bool) error {
	teacher, err := s.teacherRepo.FindByID(teacherID)
	if err != nil {
		return s.handleError(err, "Teacher not found", 400)
	}

	if teacher.IsHomeroom == isHomeroom {
		return s.handleError(err, "Homeroom status already updated", 400)
	}

	teacher.IsHomeroom = isHomeroom
	err = s.teacherRepo.Update(teacher)
	return s.handleError(err, "Failed to update teacher", 500)
}
