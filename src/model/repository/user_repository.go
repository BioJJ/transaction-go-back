package repository

import (
	"github.com/BioJJ/transaction-go-back/src/config/rest_err"
	"github.com/BioJJ/transaction-go-back/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByID(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr

	DeleteUser(
		userId string,
	) *rest_err.RestErr
}
