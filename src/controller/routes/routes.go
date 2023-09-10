package routes

import (
	"github.com/BioJJ/transaction-go-back/src/controller"
	"github.com/BioJJ/transaction-go-back/src/model"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/users/find/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/users/:userEmail/email", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:userId", model.VerifyTokenMiddleware, userController.UpdateUser)
	r.DELETE("/users/:userId", model.VerifyTokenMiddleware, userController.DeleteUser)

	r.POST("/login", userController.LoginUserServices)

}
