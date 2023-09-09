package model

import (
	"github.com/BioJJ/transaction-go-back/src/config/logger"
	"github.com/BioJJ/transaction-go-back/src/config/rest_err"
	"go.uber.org/zap"
)

func (*UserDomain) UpdateUser(string) *rest_err.RestErr {

	logger.Info("Init UpdateUser Model", zap.String("journey", "UpdateUser"))

	return nil
}
