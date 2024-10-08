package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"crud/models"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "root"
const DB_NAME = "my_db"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var dataBaseInstance *gorm.DB

func InitDb() *gorm.DB {

	dataBaseInstance = connectDB()

	dataBaseInstance.AutoMigrate(
		&models.User{}, 
		&models.Country{}
	)

	return dataBaseInstance
}

func connectDB() (*gorm.DB) {
	var err error
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME) // DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	
	fmt.Println("dsn : ", connectionString)

	databaseConnection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		// panic("Failed to connect to database");
		return nil
	}

	return databaseConnection
}



// Non-Destructive Column Operations (These operations modify the schema without affecting the existing data in the table.)
// 1. Adding New Columns.
// 2. Changing Column Attributes i.e not null, autoIncrement, etc.
// 3. Adding Indexes

// Destructive Column Operations
// 1. Dropping Columns (column dropped will lose its data)
// 2. Renaming Columns
// 3. Changing Column Data Types

