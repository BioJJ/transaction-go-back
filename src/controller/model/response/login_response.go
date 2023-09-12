package response

type LoginResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	DateBirth string `json:"dateBirth"`
	Token     string `json:"token"`
}
