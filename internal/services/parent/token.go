package service

import "github.com/Magetan-Boyz/Backend/internal/services"

type ParentTokenService interface {
	GetParentIDByUserID(userID string) (string, error)
	GetStudentIDByParentID(parentID string) (string, error)
}

func (s *parentService) GetParentIDByUserID(userID string) (string, error) {
	parentID, err := s.tokenRepo.GetParentIDByUserID(userID)
	if err != nil {
		return "", services.HandleError(err, "Failed to fetch parent", 500)
	}

	return parentID, nil
}

func (s *parentService) GetStudentIDByParentID(parentID string) (string, error) {
	studentID, err := s.parentRepo.GetStudentIDByParentID(parentID)
	if err != nil {
		return "", services.HandleError(err, "Failed to fetch student", 500)
	}

	return studentID, nil
}
