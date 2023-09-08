package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BioJJ/transaction-go-back/src/controller/routes"
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

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
