package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/naoki/english/app/configs"
	"github.com/naoki/english/app/controllers"
	"github.com/naoki/english/app/database"
	"github.com/naoki/english/app/middlewares"
	"github.com/naoki/english/app/repository"
	"github.com/naoki/english/app/service"
	"github.com/sirupsen/logrus"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (server *Server) Run() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	/* @Setup gin.Engin */
	// routes.Routes(server.Router)
	/* @Setup middlewares */

	/* @Setup godotenv */
	var err error
	err = godotenv.Load()
	if err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	config := configs.GetConfig()

	fmt.Println(config.Postgres.Dialect(), config.Postgres.GetPostgresConnectionInfo())

	/* @Connects Database */
	database.Connection()
	defer database.Close()

	/* @Connects Repository */
	repos := repository.NewRepository(database.DB)
	services := service.NewService(repos)
	handlers := controllers.NewHandler(services)

	// handlers.Routes()

	// @Connects Server
	fmt.Printf("connect to http://%s:%s/ for Gin", config.Host, config.Port)
	port := fmt.Sprintf(":%s", config.Port)
	servers := handlers.Routes()

	servers.Use(middlewares.CORSMiddleware())
	servers.Use(gin.Logger())
	servers.Use(gin.Recovery())
	servers.Run(port)
}
