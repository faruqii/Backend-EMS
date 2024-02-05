package entities

type Parent struct {
	User
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Occupation  string    `json:"occupation"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Children    []Student `json:"children"`
}
