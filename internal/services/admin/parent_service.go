package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
	"golang.org/x/crypto/bcrypt"
)

type AdminParentService interface {
	CreateParent(parent *entities.Parent) error
	AssignParentToStudent(parentID, studentID string) error
	GetAll() ([]entities.Parent, error)
}

func (s *adminService) CreateParent(parent *entities.Parent) error {
	_, err := s.userRepo.FindByUsername(parent.User.Username)
	if err == nil {
		return services.HandleError(err, "Username already exist", 400)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(parent.User.Password), bcrypt.MinCost)
	if err != nil {
		return services.HandleError(err, "Failed to hash password", 500)
	}

	parent.User.Password = string(hashedPassword)

	err = s.parentRepo.Insert(parent)
	if err != nil {
		return services.HandleError(err, "Failed to create parent", 500)
	}

	err = s.roleRepo.AssignUserRole(parent.User.ID, "parent")
	return services.HandleError(err, "Failed to assign role to parent", 500)

}

func (s *adminService) AssignParentToStudent(parentID, studentID string) error {
	// check if parent exist
	parent, err := s.parentRepo.FindById(parentID)
	if err != nil {
		return services.HandleError(err, "Parent not found", 404)
	}

	// check if student exist
	student, err := s.studentRepo.FindById(studentID)
	if err != nil {
		return services.HandleError(err, "Student not found", 404)
	}

	// check if parent already assigned to student
	_, err = s.parentRepo.FindByParentAndStudent(parentID, studentID)
	if err == nil {
		return services.HandleError(err, "Parent already assigned to student", 400)
	}

	parstud := &entities.ParentStudent{
		ParentID:  parent.ID,
		StudentID: student.ID,
	}

	err = s.parentRepo.InsertParentToStudent(parstud)
	if err != nil {
		return services.HandleError(err, "Failed to assign parent to student", 500)
	}

	return nil
}

func (s *adminService) GetAll() ([]entities.Parent, error) {
	parents, err := s.parentRepo.GetAll()
	if err != nil {
		return nil, services.HandleError(err, "Failed to get parents", 500)
	}

	return parents, nil
}
