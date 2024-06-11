package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
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
	UpdateTeacherIsCouncelorStatus(teacherID string, isCouncelor bool) error
}

func (s *adminService) CreateTeacher(teacher *entities.Teacher) error {
	_, err := s.userRepo.FindByUsername(teacher.User.Username)
	if err == nil {
		return services.HandleError(err, "Username already exist", 400)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(teacher.User.Password), bcrypt.MinCost)
	if err != nil {
		return services.HandleError(err, "Failed to hash password", 500)
	}

	teacher.User.Password = string(hashedPassword)

	err = s.teacherRepo.Create(teacher)
	if err != nil {
		return services.HandleError(err, "Failed to create teacher", 500)
	}

	err = s.roleRepo.AssignUserRole(teacher.User.ID, "teacher")
	return services.HandleError(err, "Failed to assign role to teacher", 500)
}

func (s *adminService) GetAllTeacher() ([]entities.Teacher, error) {
	teachers, err := s.teacherRepo.GetAll()
	return teachers, services.HandleError(err, "Failed to fetch teachers", 500)
}

func (s *adminService) AssignTeacherToSubject(teacherID, SubjectID string) error {
	isAssigned, err := s.subjectRepo.IsTeacherAssignedToSubject(teacherID, SubjectID)
	if err != nil {
		return services.HandleError(err, "Failed to check if teacher is assigned to subject", 500)
	}
	if isAssigned {
		return services.HandleError(err, "Teacher already assigned to subject", 400)
	}

	err = s.subjectRepo.AssignTeacherToSubject(teacherID, SubjectID)
	return services.HandleError(err, "Failed to assign teacher to subject", 500)
}

func (s *adminService) FindTeacherByID(id string) (*entities.Teacher, error) {
	teacher, err := s.teacherRepo.FindByID(id)
	return teacher, services.HandleError(err, "Failed to fetch teacher", 500)
}

func (s *adminService) GetTeachersBySubjectID(subjectID string) ([]dto.TeacherSubjectResponse, error) {
	teacherSubjects, err := s.subjectRepo.GetTeachersBySubjectID(subjectID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch teachers", 500)
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
	// check cache first
	if cachedSubjects, found := s.cache.Get(teacherID); found {
		return *cachedSubjects.(*[]dto.TeacherSubjectsResponse), nil
	}

	teacherSubjects, err := s.subjectRepo.GetTeacherSubjects(teacherID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch teacher subjects", 500)
	}

	teacherSubjectsMap := make(map[string]dto.TeacherSubjectsResponse)
	for _, ts := range teacherSubjects {
		subject := dto.SubjectResponse{
			ID:          ts.Subject.ID,
			Name:        ts.Subject.Name,
			Description: ts.Subject.Description,
			Semester:    ts.Subject.Semester,
		}

		if _, exists := teacherSubjectsMap[ts.Teacher.ID]; !exists {
			teacherSubjectsMap[ts.Teacher.ID] = dto.TeacherSubjectsResponse{
				TeacherID:   ts.Teacher.ID,
				TeacherName: ts.Teacher.Name,
				Subjects:    []dto.SubjectResponse{subject},
			}
		} else {
			tsr := teacherSubjectsMap[ts.Teacher.ID]
			tsr.Subjects = append(tsr.Subjects, subject)
			teacherSubjectsMap[ts.Teacher.ID] = tsr
		}
	}

	var subjects []dto.TeacherSubjectsResponse
	for _, tsr := range teacherSubjectsMap {
		subjects = append(subjects, tsr)
	}

	return subjects, nil
}

func (s *adminService) UpdateTeacherHomeroomStatus(teacherID string, isHomeroom bool) error {
	teacher, err := s.teacherRepo.FindByID(teacherID)
	if err != nil {
		return services.HandleError(err, "Teacher not found", 400)
	}

	if teacher.IsHomeroom == isHomeroom {
		return services.HandleError(err, "Homeroom status already updated", 400)
	}

	teacher.IsHomeroom = isHomeroom
	err = s.teacherRepo.Update(teacher)
	return services.HandleError(err, "Failed to update teacher", 500)
}

func (s *adminService) UpdateTeacherIsCouncelorStatus(teacherID string, isCouncelor bool) error {
	teacher, err := s.teacherRepo.FindByID(teacherID)
	if err != nil {
		return services.HandleError(err, "Teacher not found", 400)
	}

	if teacher.IsCouncelor == isCouncelor {
		return services.HandleError(err, "Councelor status already updated", 400)
	}

	teacher.IsCouncelor = isCouncelor
	err = s.teacherRepo.Update(teacher)
	return services.HandleError(err, "Failed to update teacher", 500)
}
