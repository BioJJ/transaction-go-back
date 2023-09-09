package model

import "github.com/BioJJ/transaction-go-back/src/config/rest_err"

type UserDomain struct {
	// id        string
	email     string
	password  string
	name      string
	phone     string
	dateBirth string
}

func NewUserDomain(
	email, password, name string,
	phone, dateBirth string,
) UserDomainInterface {
	return &UserDomain{
		email,
		password,
		name,
		phone,
		dateBirth,
	}
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

// func (ud *userDomain) GetID() string {
// 	return ud.id
// }

// func (ud *userDomain) SetID(id string) {
// 	ud.id = id
// }

// func (ud *userDomain) GetEmail() string {
// 	return ud.email
// }

// func (ud *userDomain) GetPassword() string {
// 	return ud.password
// }

// func (ud *userDomain) GetName() string {
// 	return ud.name
// }

// func (ud *userDomain) GetPhone() string {
// 	return ud.phone
// }

// func (ud *userDomain) GetDateBirth() string {
// 	return ud.dateBirth
// }
