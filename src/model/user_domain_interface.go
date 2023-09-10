package model

// import "github.com/BioJJ/transaction-go-back/src/config/rest_err"

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetPhone() string
	GetDateBirth() string
	GetID() string

	SetID(string)

	EncryptPassword()
	// GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(
	email, password, name string,
	phone, dateBirth string,
) UserDomainInterface {
	return &userDomain{
		email:     email,
		password:  password,
		name:      name,
		phone:     phone,
		dateBirth: dateBirth,
	}
}

// func NewUserLoginDomain(
// 	email, password string,
// ) UserDomainIntrest_err
// 	return &userDomain{
// 		email:    email,
// 		password: password,
// 	}
// }

func NewUserUpdateDomain(
	name string,
	phone, dateBirth string,
) UserDomainInterface {
	return &userDomain{
		name:      name,
		phone:     phone,
		dateBirth: dateBirth,
	}
}
