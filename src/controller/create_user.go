package controller

import (
	"fmt"

	"github.com/BioJJ/transaction-go-back/src/config/logger"
	"github.com/BioJJ/transaction-go-back/src/config/validation"
	"github.com/BioJJ/transaction-go-back/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {

	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
