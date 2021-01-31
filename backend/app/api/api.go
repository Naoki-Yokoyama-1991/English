package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/naoki/gacha/app/configs"
)

func Run() {

	/* @Setup gin.Engin */
	var router = gin.Default()

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
	db, err := gorm.Open(config.Postgres.Dialect(), config.Postgres.GetPostgresConnectionInfo())
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to postgres")
	db.LogMode(true)
	defer db.Close()

	// @Connects Server
	fmt.Printf("connect to http://%s:%s/ for Gin", config.Host, config.Port)
	port := fmt.Sprintf(":%s", config.Port)
	router.Run(port)
}
