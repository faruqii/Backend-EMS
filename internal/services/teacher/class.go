package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherClassService interface {
	GetWhereIamTeachTheClass(userID string) ([]entities.ClassSubject, error)
	IsIamHomeroomTeacher(userID string, classID string) (bool, error)
	GetAllStudents(classID string) ([]entities.Student, error)
}

func (s *teacherService) GetWhereIamTeachTheClass(userID string) ([]entities.ClassSubject, error) {
	teacherID, err := s.tokenRepo.GetTeacherIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch teacher", 500)
	}

	class, err := s.subjectRepo.GetWhereIamTeachTheClass(teacherID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch class", 500)
	}

	return class, nil
}

func (s *teacherService) IsIamHomeroomTeacher(userID string, classID string) (bool, error) {
	teacherID, err := s.tokenRepo.GetTeacherIDByUserID(userID)
	if err != nil {
		return false, services.HandleError(err, "Failed to fetch teacher", 500)
	}

	class, err := s.classRepo.IsTeacherHomeRoomTeacher(teacherID, classID)
	if err != nil {
		return false, services.HandleError(err, "Failed to fetch class", 500)
	}

	return class, nil
}

func (s *teacherService) GetAllStudents(classID string) ([]entities.Student, error) {
	students, err := s.classRepo.GetAllStudents(classID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch students", 500)
	}

	return students, nil
}
