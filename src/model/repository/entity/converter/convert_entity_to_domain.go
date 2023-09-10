package converter

import (
	"github.com/BioJJ/transaction-go-back/src/model"
	"github.com/BioJJ/transaction-go-back/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Phone,
		entity.DateBirth,
	)

	domain.SetID(entity.ID.Hex())
	return domain
}
