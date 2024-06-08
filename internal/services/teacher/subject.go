package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherSubjectService interface {
	CountStudent(classID, subjectID string) ([]dto.StudentResponse, error)
}

func (s *teacherService) CountStudent(classID, subjectID string) ([]dto.StudentResponse, error) {
	students, err := s.subjectRepo.GetStudentsByClassAndSubjectID(classID, subjectID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get students", 500)
	}

	var studentResponses []dto.StudentResponse
	for _, student := range students {
		studentResponses = append(studentResponses, dto.StudentResponse{
			ID:         student.ID,
			Name:       student.Name,
			NISN:       student.NISN,
			Address:    student.Address,
			Birthplace: student.Birthplace,
			Birthdate:  student.Birthdate,
		})
	}

	return studentResponses, nil
}
