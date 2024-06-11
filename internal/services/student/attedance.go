package service

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type StudentAttendanceService interface {
	MyAttedance(userID string) ([]entities.Atendance, error)
}

func (s *studentService) MyAttedance(userID string) ([]entities.Atendance, error) {
	studentID, err := s.tokenRepo.GetStudentIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch student", 500)
	}

	attedance, err := s.attedanceRepo.GetMyAttedance(studentID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to fetch attedance", 500)
	}

	return attedance, nil
}
