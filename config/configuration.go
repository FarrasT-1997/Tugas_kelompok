package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init_DB() {
	config := map[string]string{
		"HTTP_USERNAME": "root",
		"HTTP_PASSWORD": "aPasswordNeeded",
		"HTTP_PORT":     "3306",
		"HTTP_HOST":     "127.0.0.1",
		"HTTP_NAME":     "new_schema",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["HTTP_USERNAME"],
		config["HTTP_PASSWORD"],
		config["HTTP_HOST"],
		config["HTTP_PORT"],
		config["HTTP_NAME"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}
