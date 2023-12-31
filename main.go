package main

import (
	"context"
	"log"

	"github.com/BioJJ/transaction-go-back/src/config/database/mongodb"
	"github.com/BioJJ/transaction-go-back/src/controller"
	"github.com/BioJJ/transaction-go-back/src/controller/routes"
	"github.com/BioJJ/transaction-go-back/src/model/repository"
	"github.com/BioJJ/transaction-go-back/src/model/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept", "authorization"}
	config.AllowMethods = []string{"PUT", "GET", "POST", "DELETE", "OPTIONS"}

	router.Use(cors.New(config))

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
