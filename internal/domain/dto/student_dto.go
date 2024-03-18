package dto

type StudentRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	NISN       string `json:"nisn"`
	Address    string `json:"address"`
	Birthplace string `json:"birthplace"`
	Birthdate  string `json:"birthdate"`
}

type StudentResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	NISN       string `json:"nisn"`
	Address    string `json:"address"`
	Birthplace string `json:"birthplace"`
	Birthdate  string `json:"birthdate"`
}
