package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Ganti driver mariadb.Open() dengan mysql.Open()
	database, err := gorm.Open(mysql.Open("root:root@tcp(192.168.131.129:3306)/sisbus"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Studi{})

	DB = database
}
