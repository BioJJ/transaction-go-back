package routes

import (
	"github.com/BioJJ/transaction-go-back/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/users/find/:id", userController.FindUserById)
	r.GET("/users/:email/email", userController.FindUserByEmail)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/delete/:id", userController.DeleteUser)

}
