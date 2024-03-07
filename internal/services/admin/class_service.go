package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type AdminClassService interface {
	CreateClass(class *entities.Class) error
	AssignHomeroomTeacher(classID, teacherID string) error
	FindClassByID(id string) (*entities.Class, error)
	GetAllClass() ([]entities.Class, error)
	GetClassSchedule(classID string) ([]entities.Schedule, error)
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
		return services.HandleError(err, "Teacher not found", 400)
	}

	class, err := s.classRepo.FindByID(classID)
	if err != nil {
		return services.HandleError(err, "Class not found", 400)
	}

	if class.HomeRoomTeacherID != nil {
		return services.HandleError(err, "Class already has a homeroom teacher", 400)
	}

	if teacher.IsHomeroom {
		return services.HandleError(err, "Teacher is already a homeroom teacher", 400)
	}

	class.HomeRoomTeacherID = &teacherID
	if err := s.classRepo.Update(class); err != nil {
		return services.HandleError(err, "Failed to assign teacher as homeroom", 500)
	}

	teacher.IsHomeroom = true
	if err := s.teacherRepo.Update(teacher); err != nil {
		return services.HandleError(err, "Failed to update teacher", 500)
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
