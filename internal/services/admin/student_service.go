package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
	"golang.org/x/crypto/bcrypt"
)

type AdminStudentService interface {
	CreateStudent(student *entities.Student) error
}

func (s *adminService) CreateStudent(student *entities.Student) error {
	_, err := s.studentRepo.FindByUsername(student.User.Username)
	if err == nil {
		return services.HandleError(err, "Student already exist", 400)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.User.Password), bcrypt.MinCost)
	if err != nil {
		return services.HandleError(err, "Failed to hash password", 500)
	}

	student.User.Password = string(hashedPassword)

	err = s.studentRepo.Insert(student)
	if err != nil {
		return services.HandleError(err, "Failed to create teacher", 500)
	}

	err = s.roleRepo.AssignUserRole(student.User.ID, "student")
	return services.HandleError(err, "Failed to assign role to student", 500)
}
