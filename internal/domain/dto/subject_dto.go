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
	TeacherName string   `json:"teacher_name"`
	SubjectName []string `json:"subject_name"`
}
