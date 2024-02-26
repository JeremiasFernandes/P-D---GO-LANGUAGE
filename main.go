package main

import (
	"context"
	"log"

	_ "github.com/JeremiasFernandes/PeDGo/docs"
	"github.com/JeremiasFernandes/PeDGo/src/configuration/database/mongodb"
	"github.com/JeremiasFernandes/PeDGo/src/configuration/logger"
	"github.com/JeremiasFernandes/PeDGo/src/controller"
	"github.com/JeremiasFernandes/PeDGo/src/controller/routes"
	"github.com/JeremiasFernandes/PeDGo/src/model/repository"
	"github.com/JeremiasFernandes/PeDGo/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}

// @title PeD Go
// @version 1.0
// @description API for crud operations on users
// @host localhost:8080
// @BasePath /
// @schemes http
// @license MIT
func main() {
	logger.Info("Iniciando a aplicação")

	godotenv.Load()

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	if err != nil {
		panic(err)
	}

	userController := initDependencies(database)

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
