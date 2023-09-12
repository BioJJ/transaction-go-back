package view

import (
	"github.com/BioJJ/transaction-go-back/src/controller/model/response"
	"github.com/BioJJ/transaction-go-back/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:        userDomain.GetID(),
		Email:     userDomain.GetEmail(),
		Name:      userDomain.GetName(),
		Phone:     userDomain.GetPhone(),
		DateBirth: userDomain.GetDateBirth(),
	}
}

func ConvertDomainToResponseLogin(
	userDomain model.UserDomainInterface,
	token string,
) response.LoginResponse {
	return response.LoginResponse{
		ID:        userDomain.GetID(),
		Email:     userDomain.GetEmail(),
		Name:      userDomain.GetName(),
		Phone:     userDomain.GetPhone(),
		DateBirth: userDomain.GetDateBirth(),
		Token:     token,
	}
}
