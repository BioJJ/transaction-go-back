package model

import (
	"github.com/BioJJ/transaction-go-back/src/config/logger"
	"github.com/BioJJ/transaction-go-back/src/config/rest_err"
	"go.uber.org/zap"
)

func (*UserDomain) DeleteUser(string) *rest_err.RestErr {

	logger.Info("Init DeleteUser Model", zap.String("journey", "DeleteUser"))

	return nil
}
