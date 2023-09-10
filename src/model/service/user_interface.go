package service

import (
	"github.com/BioJJ/transaction-go-back/src/config/rest_err"
	"github.com/BioJJ/transaction-go-back/src/model"
	"github.com/BioJJ/transaction-go-back/src/model/repository"
)

type userDomainService struct {
	userRepository repository.UserRepository
}

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
