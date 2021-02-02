package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/naoki/gacha/app/configs"
	"github.com/naoki/gacha/app/models"
)

var DB *gorm.DB

func Connection() {

	config := configs.GetConfig()

	db, err := gorm.Open(config.Postgres.Dialect(), config.Postgres.GetPostgresConnectionInfo())
	if err != nil {
		panic(err)
	}

	DB = db
	fmt.Println("connect to postgres")
	DB.LogMode(true)

	DB.AutoMigrate(&models.User{})
}

func Close() {
	if DB == nil {
		return
	}
	DB.Close()
}
