package entities

type Teacher struct {
	User
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	IsHomeroom bool      `json:"isHomeroom"`
	Subjects   []Subject `json:"subject" gorm:"many2many:teacher_subjects"`
}
