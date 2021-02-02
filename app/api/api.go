package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/naoki/gacha/app/configs"
	"github.com/naoki/gacha/app/database"
	"github.com/naoki/gacha/app/routes"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (server *Server) Run() {
	/* @Setup gin.Engin */
	server.Router = gin.New()
	routes.Routes(server.Router)
	/* @Setup middlewares */
	server.Router.Use(gin.Logger())
	server.Router.Use(gin.Recovery())

	/* @Setup godotenv */
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting values")
	}
	config := configs.GetConfig()

	fmt.Println(config.Postgres.Dialect(), config.Postgres.GetPostgresConnectionInfo())

	/* @Connects Database */
	database.Connection()
	defer database.Close()

	// @Connects Server
	fmt.Printf("connect to http://%s:%s/ for Gin", config.Host, config.Port)
	port := fmt.Sprintf(":%s", config.Port)
	server.Router.Run(port)
}
