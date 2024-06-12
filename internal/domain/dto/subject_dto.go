package dto

type SubjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Semester    string `json:"semester"`
}

type SubjectResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Semester    string `json:"semester"`
}

type TeacherSubjectRequest struct {
	TeacherID string `json:"teacher_id"`
	SubjectID string `json:"subject_id"`
}

type TeacherSubjectResponse struct {
	SubjectName string `json:"subject_name"`
	TeacherName string `json:"teacher_name"`
}

type TeacherSubjectsResponse struct {
	TeacherID   string            `json:"teacher_id"`
	TeacherName string            `json:"teacher_name"`
	Subjects    []SubjectResponse `json:"subject"`
}

type AssignSubjectToClassRequest struct {
	SubjectID string `json:"subject_id"`
	TeacherID string `json:"teacher_id"`
	ClassID   string `json:"class_id"`
}

type ClassSubjectResponse struct {
	ClassName   string `json:"class_name"`
	SubjectID   string `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	TeacherName string `json:"teacher_name"`
}

type SubjectMattterRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

type SubjectMattterResponse struct {
	ID          string `json:"id"`
	Subject     string `json:"subject"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
