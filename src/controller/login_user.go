package controller

import (
	"net/http"

	"github.com/BioJJ/transaction-go-back/src/config/logger"
	"github.com/BioJJ/transaction-go-back/src/config/validation"
	"github.com/BioJJ/transaction-go-back/src/controller/model/request"
	"github.com/BioJJ/transaction-go-back/src/model"
	"github.com/BioJJ/transaction-go-back/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUserServices(c *gin.Context) {
	logger.Info("Init loginUser controller",
		zap.String("journey", "loginUser"),
	)
	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "loginUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)
	domainResult, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error(
			"Error trying to call loginUser service",
			err,
			zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"loginUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "loginUser"))

	// c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}
