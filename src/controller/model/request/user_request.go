package request

type UserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	Name      string `json:"name" binding:"required,min=4,max=100"`
	Phone     string `json:"phone" binding:"required,min=1,max=140"`
	DateBirth string `json:"dateBirth" binding:"required,min=1,max=140"`
}

type UserUpdateRequest struct {
	Name      string `json:"name" binding:"omitempty,min=4,max=100"`
	Phone     string `json:"phone" binding:"omitempty,min=1,max=140"`
	DateBirth string `json:"dateBirth" binding:"omitempty,min=1,max=140"`
}
