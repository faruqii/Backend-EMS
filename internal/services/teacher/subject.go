package services

import (
	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherSubjectService interface {
	CountStudent(classID, subjectID string) ([]dto.StudentResponse, error)
	GetMySubjects(userID string) ([]dto.TeacherSubjectsResponse, error)
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

func (s *teacherService) GetMySubjects(userID string) ([]dto.TeacherSubjectsResponse, error) {
	teacherID, err := s.tokenRepo.GetTeacherIDByUserID(userID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get teacher ID", 500)
	}

	teacherSubjects, err := s.subjectRepo.GetTeacherSubjects(teacherID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get teacher subjects", 500)
	}

	teacherSubjectsMap := make(map[string]dto.TeacherSubjectsResponse)
	for _, ts := range teacherSubjects {
		subject := dto.SubjectResponse{
			ID:          ts.Subject.ID,
			Name:        ts.Subject.Name,
			Description: ts.Subject.Description,
			Semester:    ts.Subject.Semester,
		}

		if _, exists := teacherSubjectsMap[ts.Teacher.ID]; !exists {
			teacherSubjectsMap[ts.Teacher.ID] = dto.TeacherSubjectsResponse{
				TeacherID:   ts.Teacher.ID,
				TeacherName: ts.Teacher.Name,
				Subjects:    []dto.SubjectResponse{subject},
			}
		} else {
			tsr := teacherSubjectsMap[ts.Teacher.ID]
			tsr.Subjects = append(tsr.Subjects, subject)
			teacherSubjectsMap[ts.Teacher.ID] = tsr
		}
	}

	var subjects []dto.TeacherSubjectsResponse
	for _, tsr := range teacherSubjectsMap {
		subjects = append(subjects, tsr)
	}

	return subjects, nil

}
