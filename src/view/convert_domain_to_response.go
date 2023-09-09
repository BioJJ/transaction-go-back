package view

import (
	"github.com/BioJJ/transaction-go-back/src/controller/model/response"
	"github.com/BioJJ/transaction-go-back/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		// ID:    userDomain.GetID(),
		ID:        "",
		Email:     userDomain.GetEmail(),
		Name:      userDomain.GetName(),
		Phone:     userDomain.GetPhone(),
		DateBirth: userDomain.GetDateBirth(),
	}
}
