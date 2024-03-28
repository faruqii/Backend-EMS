package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
	"golang.org/x/crypto/bcrypt"
)

type AdminStudentService interface {
	CreateStudent(student *entities.Student) error
	GetAllStudents() ([]entities.Student, error)
	InsertStudentToClass(studentID, classID string) (*entities.Student, error)
}

func (s *adminService) CreateStudent(student *entities.Student) error {
	_, err := s.studentRepo.FindByUsername(student.User.Username)
	if err == nil {
		return services.HandleError(err, "Student already exist", 400)
	}

	// chek if nisn already exist
	_, err = s.studentRepo.FindByNISN(student.NISN)
	if err == nil {
		return services.HandleError(err, "NISN already exist", 400)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.User.Password), bcrypt.MinCost)
	if err != nil {
		return services.HandleError(err, "Failed to hash password", 500)
	}

	student.User.Password = string(hashedPassword)

	err = s.studentRepo.Insert(student)
	if err != nil {
		return services.HandleError(err, "Failed to create student", 500)
	}

	err = s.roleRepo.AssignUserRole(student.User.ID, "student")
	return services.HandleError(err, "Failed to assign role to student", 500)
}

func (s *adminService) GetAllStudents() ([]entities.Student, error) {
	students, err := s.studentRepo.GetAllStudents()
	if err != nil {
		return nil, services.HandleError(err, "Failed to get students", 500)
	}

	return students, nil
}

func (s *adminService) InsertStudentToClass(studentID, classID string) (*entities.Student, error) {
	// check if student already in class
	isStudentInClass, err := s.studentRepo.IsStudentAlreadyInClass(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to check student in class", 500)
	}

	if isStudentInClass {
		return nil, services.HandleError(err, "Student already in class", 400)
	}

	student, err := s.studentRepo.InsertStudentToClass(studentID, classID)

	if err != nil {
		return nil, services.HandleError(err, "Failed to insert student to class", 500)
	}

	return student, nil

}
