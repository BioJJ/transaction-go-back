package service

import (
	"github.com/BioJJ/transaction-go-back/src/config/logger"
	"github.com/BioJJ/transaction-go-back/src/config/rest_err"
	"github.com/BioJJ/transaction-go-back/src/model"
	"go.uber.org/zap"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init FindUser Model", zap.String("journey", "FindUser"))

	return nil, nil
}
