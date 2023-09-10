package routes

import (
	"github.com/BioJJ/transaction-go-back/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/users/find/:userId", userController.FindUserByID)
	r.GET("/users/:userEmail/email", userController.FindUserByEmail)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:userId", userController.UpdateUser)
	r.DELETE("/users/:userId", userController.DeleteUser)

	r.POST("/login", userController.LoginUserServices)

}
