package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/naoki/english/app/configs"
	"github.com/naoki/english/app/models"
)

var DB *gorm.DB

func Connection() {

	config := configs.GetConfig()

	db, err := gorm.Open(config.Postgres.Dialect(), config.Postgres.GetPostgresConnectionInfo())
	if err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		return
	}

	DB = db
	fmt.Println("connect to postgres")
	DB.LogMode(true)

	DB.AutoMigrate(&models.Phrase{})
}

func Close() {
	if DB == nil {
		return
	}
	DB.Close()
}
