package database

import (
	"fmt"

	"github.com/jagadwp/gin-crud-web/config"
	"github.com/jagadwp/gin-crud-web/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB
var err error

func DatabaseInit() {
	//Connect to database, exit when errored
	database, err = gorm.Open("mysql", "root:@/berat?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	fmt.Println("DB Connected")

	//If set true then print all executed queries to the console
	database.LogMode(config.DbDebug)

	//Migrate database
	initMigrate()
}

func DB() *gorm.DB {
	return database
}

func initMigrate() {
	if err := database.AutoMigrate(&model.Berat{}).Error; err != nil {
		panic("Failed migrating database")
	}
}
