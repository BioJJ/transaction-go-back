package model

type userDomain struct {
	// id        string
	email     string
	password  string
	name      string
	phone     string
	dateBirth string
}

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetPhone() string
	GetDateBirth() string
	// GetID() string
	// SetID(string)

	EncryptPassword()
	// GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(
	email, password, name string,
	phone, dateBirth string,
) UserDomainInterface {
	return &userDomain{
		email,
		password,
		name,
		phone,
		dateBirth,
	}
}

// func (ud *userDomain) GetID() string {
// 	return ud.id
// }

// func (ud *userDomain) SetID(id string) {
// 	ud.id = id
// }

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetPhone() string {
	return ud.phone
}

func (ud *userDomain) GetDateBirth() string {
	return ud.dateBirth
}
