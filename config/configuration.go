package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init_DB() {
	var e error
	DB, e = gorm.Open(mysql.Open(ConnectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}
