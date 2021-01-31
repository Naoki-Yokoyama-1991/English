package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/naoki/gacha/app/configs"
	"github.com/naoki/gacha/app/routes"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (server *Server) Run() {

	/* @Setup gin.Engin */
	server.Router = gin.Default()
	routes.Routes(server.Router)
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
	server.DB, err = gorm.Open(config.Postgres.Dialect(), config.Postgres.GetPostgresConnectionInfo())
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to postgres")
	server.DB.LogMode(true)
	defer server.DB.Close()

	// @Connects Server
	fmt.Printf("connect to http://%s:%s/ for Gin", config.Host, config.Port)
	port := fmt.Sprintf(":%s", config.Port)
	server.Router.Run(port)
}
