package entities

type Parent struct {
	User
	Name        string     `json:"name"`
	Address     string     `json:"address"`
	Occupation  string     `json:"occupation"`
	PhoneNumber string     `json:"phone_number"`
	Email       string     `json:"email"`
	Students    []*Student `json:"students"` // One parent can have many children that study in the same school
}