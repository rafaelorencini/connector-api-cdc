package database

import (
	"connector-api-cdc/domain"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "root"
const DB_NAME = "my_db"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&domain.Connector{}, &domain.Tables{})

	// Add table suffix when creating tables
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&repository.User{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
