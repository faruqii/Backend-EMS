package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentProfile interface {
	MyProfile(userID string) (*dto.StudentResponse, error)
}

func (s *studentService) MyProfile(userID string) (*dto.StudentResponse, error) {
	student, err := s.studentRepo.GetStudentByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get student", 500)
	}

	return &dto.StudentResponse{
		ID:          student.ID,
		Name:        student.Name,
		NISN:        student.NISN,
		Address:     student.Address,
		Birthplace:  student.Birthplace,
		Birthdate:   student.Birthdate,
		Gender:      student.Gender,
		Province:    student.Province,
		City:        student.City,
		BloodType:   student.BloodType,
		Religion:    student.Religion,
		Phone:       student.Phone,
		ParentPhone: student.ParentPhone,
		Email:       student.Email,
	}, nil
}
