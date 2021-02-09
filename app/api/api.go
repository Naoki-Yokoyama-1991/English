package api

import (
	"fmt"

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

func Run() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	/* @Setup gin.Engin */

	/* @Setup godotenv */
	var err error
	err = godotenv.Load()
	if err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	/* @Get env*/
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
	server := handlers.Handler()

	/* @Setup middlewares */
	server.Use(middlewares.CORSMiddleware())

	// @Connects Server
	fmt.Printf("connect to http://%s:%s/ for Gin", config.Host, config.Port)
	port := fmt.Sprintf(":%s", config.Port)
	server.Run(port)
}
