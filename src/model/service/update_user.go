package service

import (
	"github.com/BioJJ/transaction-go-back/src/config/logger"
	"github.com/BioJJ/transaction-go-back/src/config/rest_err"
	"github.com/BioJJ/transaction-go-back/src/model"
	"go.uber.org/zap"
)

func (*userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("Init UpdateUser Model", zap.String("journey", "UpdateUser"))

	return nil
}
