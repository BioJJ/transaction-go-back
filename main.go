package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BioJJ/transaction-go-back/src/controller"
	"github.com/BioJJ/transaction-go-back/src/controller/routes"
	"github.com/BioJJ/transaction-go-back/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("variavel de ambiente ==> ", os.Getenv("TEST"))

	router := gin.Default()

	service := service.NewUserDomainService()

	userController := controller.NewUserControllerInterface(service)

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
