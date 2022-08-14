package database

import (
	"fmt"

	"github.com/bonggar/gorestapi/config"
	"github.com/bonggar/gorestapi/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//MySqlDBConnect : Create Connection to database
func MySqlDBConnect() {
	//Connect to database, exit when errored
	db, err := gorm.Open("mysql", "root:@/berat?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	fmt.Println("DB Connected")

	//If set true then print all executed queries to the console
	db.LogMode(config.DbDebug)

	//Migrate database
	MySqlMigrate(db)
}

//MySqlMigrate : do auto migration
func MySqlMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&model.Berat{}).Error; err != nil {
		panic("Failed migrating database")
	}
}
