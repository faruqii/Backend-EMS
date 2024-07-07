package services

import (
	"fmt"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/services"
)

type TeacherSubjectService interface {
	CountStudent(classID, subjectID string) ([]dto.StudentResponse, error)
	GetMySubjects(userID string) ([]dto.TeacherSubjectsResponse, error)
	CreateSubjectMatter(subjectMatter *entities.SubjectMattter) error
	GetSubjectMatterBySubjectID(subjectID string) ([]entities.SubjectMattter, error)
	GetDetailSubjectMatter(subjectMatterID string) (*entities.SubjectMattter, error)
}

func (s *teacherService) CountStudent(classID, subjectID string) ([]dto.StudentResponse, error) {
	students, err := s.subjectRepo.GetStudentsByClassAndSubjectID(classID, subjectID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get students", 500)
	}

	var studentResponses []dto.StudentResponse
	for _, student := range students {
		studentResponses = append(studentResponses, dto.StudentResponse{
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

func (s *teacherService) CreateSubjectMatter(subjectMatter *entities.SubjectMattter) error {
    fmt.Println("Service: Entering CreateSubjectMatter")
    if err := s.subjectRepo.CreateSubjectMatter(subjectMatter); err != nil {
        fmt.Println("Service: Error creating subject matter:", err)
        return services.HandleError(err, "Failed to create subject matter", 500)
    }
    fmt.Println("Service: Successfully created subject matter")
    return nil
}


func (s *teacherService) GetSubjectMatterBySubjectID(subjectID string) ([]entities.SubjectMattter, error) {
	subjectMatters, err := s.subjectRepo.GetSubjectMatterBySubjectID(subjectID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get subject matters", 500)
	}
	return subjectMatters, nil
}

func (s *teacherService) GetDetailSubjectMatter(subjectMatterID string) (*entities.SubjectMattter, error) {
	subjectMatter, err := s.subjectRepo.GetDetailSubjectMatter(subjectMatterID)
	if err != nil {
		return nil, services.HandleError(err, "Failed to get subject matter", 500)
	}
	return subjectMatter, nil
}
