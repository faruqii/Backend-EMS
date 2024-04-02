package entities

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Grade struct {
	ID                  string          `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	StudentID           string          `json:"student_id"`
	Student             Student         `json:"student" gorm:"foreignKey:StudentID"`
	SubjectID           string          `json:"subject_id"`
	Subject             Subject         `json:"subject" gorm:"foreignKey:SubjectID"`
	TeacherID           string          `json:"teacher_id"`
	Teacher             Teacher         `json:"teacher" gorm:"foreignKey:TeacherID"`
	Semester            int             `json:"semester"`
	FormativeScores     pq.Float32Array `json:"formative_scores" gorm:"type:float[]"`   // Formative scores
	SummativeScores     pq.Float32Array `json:"summative_scores" gorm:"type:float[]"`   // Summative scores
	ProjectScores       pq.Float32Array `json:"project_scores" gorm:"type:float[]"`     // Project scores
	NonAssessmentScores pq.Float32Array `json:"non_assessment" gorm:"type:float[]"`     // Non assessment scores
	TestAssessment      pq.Float32Array `json:"test_assessment" gorm:"type:float[]"`    // Test assessment
	CompetencyScores    pq.Float32Array `json:"competency_scores" gorm:"type:float[]"`  // Competency scores
	PortfolioScores     pq.Float32Array `json:"portfolio_scores" gorm:"type:float[]"`   // Portfolio scores
	PerformanceScores   pq.Float32Array `json:"performance_scores" gorm:"type:float[]"` // Performance scores
	ProductScores       pq.Float32Array `json:"product_scores" gorm:"type:float[]"`     // Product scores
	AtendanceScores     pq.Float32Array `json:"atendance_scores" gorm:"type:float[]"`   // Atendance scores
	FinalGrade          float32         `json:"final_grade" gorm:"type:float[]"`        // Final grade
}

func (g *Grade) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = uuid.NewString()
	return nil
}
