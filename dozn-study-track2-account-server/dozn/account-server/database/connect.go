package database

import (
	"fmt"
	"strconv"
	"dozn/account-server/models"
	"gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config models.Config) {
	var err error
	p := config.DbPort
	port, err := strconv.ParseUint(p, 10, 32)
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?allowNativePasswords=true",
			config.DbUser,
			config.DbPassword,
			config.DbHost,
			port,
			config.DbName)

	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		fmt.Println("Failed to connect database")
	}

	fmt.Println("Connection Opend to Database")

	DB.AutoMigrate(&models.Account{})
	fmt.Println("Database Migrated")
}