package converter

import (
	"github.com/BioJJ/transaction-go-back/src/model"
	"github.com/BioJJ/transaction-go-back/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:     domain.GetEmail(),
		Password:  domain.GetPassword(),
		Name:      domain.GetName(),
		Phone:     domain.GetPhone(),
		DateBirth: domain.GetDateBirth(),
	}
}
