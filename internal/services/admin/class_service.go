package services

import (
	"errors"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type AdminClassService interface {
	CreateClass(class *entities.Class) error
	AssignHomeroomTeacher(classID, teacherID string) error
	FindClassByID(id string) (*entities.Class, error)
	GetAllClass() ([]entities.Class, error)
	GetClassSchedule(classID string) ([]entities.Schedule, error)
	GetAllStudentsBelongToClass(studentID string) ([]entities.Student, error)
	ClassExists(classID string) (bool, error)
	AssignSubjectToClass(subjectID, teacherID, classID string) error
}

func (s *adminService) CreateClass(class *entities.Class) error {
	_, err := s.classRepo.FindByName(class.Name)
	if err == nil {
		return services.HandleError(err, "Class already exist", 400)
	}

	err = s.classRepo.Insert(class)
	return services.HandleError(err, "Failed to create class", 500)
}

func (s *adminService) AssignHomeroomTeacher(classID, teacherID string) error {
	teacher, err := s.teacherRepo.FindByID(teacherID)
	if err != nil {
		return services.HandleError(errors.New("teacher not found"), "Teacher not found", 400)
	}

	class, err := s.classRepo.FindByID(classID)
	if err != nil {
		return services.HandleError(errors.New("class not found"), "Class not found", 400)
	}

	if class.HomeRoomTeacherID != nil {
		return services.HandleError(errors.New("homeroom teacher already assigned"), "Homeroom teacher already assigned", 400)
	}

	if teacher.IsHomeroom {
		return services.HandleError(errors.New("teacher is already homeroom"), "Teacher is already homeroom Teacher", 400)
	}

	class.HomeRoomTeacherID = &teacherID
	if err := s.classRepo.Update(class); err != nil {
		return services.HandleError(errors.New("failed to update class"), "Failed to update class", 500)
	}

	teacher.IsHomeroom = true
	if err := s.teacherRepo.Update(teacher); err != nil {
		return services.HandleError(errors.New("failed to update teacher"), "Failed to update teacher", 500)
	}

	return nil
}

func (s *adminService) FindClassByID(id string) (*entities.Class, error) {
	class, err := s.classRepo.FindByID(id)
	return class, services.HandleError(err, "Failed to fetch class", 500)
}

func (s *adminService) GetAllClass() ([]entities.Class, error) {
	classes, err := s.classRepo.GetAll()
	return classes, services.HandleError(err, "Failed to fetch classes", 500)
}

func (s *adminService) GetClassSchedule(classID string) ([]entities.Schedule, error) {
	schedules, err := s.scheduleRepo.FindByClassID(classID)
	return schedules, services.HandleError(err, "Failed to fetch class schedule", 500)
}

func (s *adminService) GetAllStudentsBelongToClass(classID string) ([]entities.Student, error) {
	students, err := s.classRepo.GetAllStudents(classID)
	return students, services.HandleError(err, "Failed to fetch students", 500)
}

func (s *adminService) ClassExists(classID string) (bool, error) {
	exists, err := s.classRepo.ClassExists(classID)
	return exists, services.HandleError(err, "Failed to check class existence", 500)
}

func (s *adminService) AssignSubjectToClass(subjectID, teacherID, classID string) error {
	_, err := s.teacherRepo.FindByID(teacherID)
	if err != nil {
		return services.HandleError(err, "Teacher not found", 400)
	}

	_, err = s.classRepo.FindByID(classID)
	if err != nil {
		return services.HandleError(err, "Class not found", 400)
	}

	_, err = s.subjectRepo.FindByID(subjectID)
	if err != nil {
		return services.HandleError(err, "Subject not found", 400)
	}

	isAssigned, err := s.subjectRepo.IsTeacherAssignedToSubject(teacherID, subjectID)
	if err != nil {
		return services.HandleError(err, "Failed to check teacher assignment", 500)
	}

	if !isAssigned {
		return services.HandleError(err, "Teacher is not assigned to subject", 400)
	}

	_, err = s.subjectRepo.AssignSubjectToClass(subjectID, teacherID, classID)
	return services.HandleError(err, "Failed to assign subject to class", 500)

}
