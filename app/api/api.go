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
	/* @Setup gin.Engin */
	server.Router = gin.New()
	// routes.Routes(server.Router)
	/* @Setup middlewares */
	server.Router.Use(middlewares.CORSMiddleware())
	server.Router.Use(gin.Logger())
	server.Router.Use(gin.Recovery())

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
	_ = controllers.NewHandler(services)
	// handlers.Routes()

	// @Connects Server
	fmt.Printf("connect to http://%s:%s/ for Gin", config.Host, config.Port)
	port := fmt.Sprintf(":%s", config.Port)
	server.Router.Run(port)
}
