package dto

type StudentRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	NISN        string `json:"nisn"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	Birthplace  string `json:"birthplace"`
	Birthdate   string `json:"birthdate"`
	Province    string `json:"province"`
	City        string `json:"city"`
	BloodType   string `json:"blood_type"`
	Religion    string `json:"religion"`
	Phone       string `json:"phone"`
	ParentPhone string `json:"parent_phone"`
	Email       string `json:"email"`
}

type StudentResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	NISN        string `json:"nisn"`
	Address     string `json:"address"`
	Birthplace  string `json:"birthplace"`
	Birthdate   string `json:"birthdate"`
	Gender      string `json:"gender"`
	Province    string `json:"province"`
	City        string `json:"city"`
	BloodType   string `json:"blood_type"`
	Religion    string `json:"religion"`
	Phone       string `json:"phone"`
	ParentPhone string `json:"parent_phone"`
	Email       string `json:"email"`
}

type InsertStudentToClass struct {
	StudentID string `json:"student_id"`
}

type StudentClassResponse struct {
	Name string `json:"name"`
}
