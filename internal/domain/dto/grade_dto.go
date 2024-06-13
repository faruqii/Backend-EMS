package dto

type GradeRequest struct {
	StudentID       string  `json:"student_id"`
	Semester        int     `json:"semester"`
	AcademicYear    string  `json:"academic_year"`
	FormativeScores float32 `json:"formative_scores"`
	SummativeScores float32 `json:"summative_scores"`
	ProjectScores   float32 `json:"project_scores"`
	FinalGrade      float32 `json:"final_grade"`
}

type GradeResponse struct {
	ID              string  `json:"id"`
	StudentID       string  `json:"student_id"`
	Student         string  `json:"student"`
	SubjectID       string  `json:"subject_id"`
	Subject         string  `json:"subject"`
	TeacherID       string  `json:"teacher_id"`
	Teacher         string  `json:"teacher"`
	Semester        int     `json:"semester"`
	AcademicYear    string  `json:"academic_year"`
	FormativeScores float32 `json:"formative_scores"`
	SummativeScores float32 `json:"summative_scores"`
	ProjectScores   float32 `json:"project_scores"`
	FinalGrade      float32 `json:"final_grade"`
}
