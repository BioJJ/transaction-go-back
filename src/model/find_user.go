package model

import (
	"github.com/BioJJ/transaction-go-back/src/config/logger"
	"github.com/BioJJ/transaction-go-back/src/config/rest_err"
	"go.uber.org/zap"
)

func (*UserDomain) FindUser(string) (*UserDomain, *rest_err.RestErr) {

	logger.Info("Init FindUser Model", zap.String("journey", "FindUser"))

	return nil, nil
}
