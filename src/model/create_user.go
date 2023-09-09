package model

import (
	"fmt"

	"github.com/BioJJ/transaction-go-back/src/config/logger"
	"github.com/BioJJ/transaction-go-back/src/config/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init createUser Model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)
	return nil
}
