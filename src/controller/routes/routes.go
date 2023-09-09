package routes

import (
	"github.com/BioJJ/transaction-go-back/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users/find/:id", controller.FindUserById)
	r.GET("/users/:email/email", controller.FindUserByEmail)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/delete/:id", controller.DeleteUser)

}
