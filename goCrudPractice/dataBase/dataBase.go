package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"crud/database/models"
)

var dataBaseInstance *gorm.DB

// InitDb initializes the database connection and performs auto-migrations
func InitDb() *gorm.DB {
	var err error
	dataBaseInstance, err = connectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return nil
	}

	// Perform auto-migrations
	err = performMigrations()
	if err != nil {
		log.Fatalf("Could not perform migrations: %v", err)
	}

	return dataBaseInstance
}

// connectDB establishes a connection to the database and configures the connection pool
func connectDB() (*gorm.DB, error) {
	// Load database connection details from environment variables
	dbUsername := getEnv("DB_USERNAME", "root")
	dbPassword := getEnv("DB_PASSWORD", "")
	dbName := getEnv("DB_NAME", "my_db")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	fmt.Println("DSN: ", connectionString)

	databaseConnection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable logging for debugging
	})

	if err != nil {
		return nil, err
	}

	// Configure connection pool
	sqlDB, err := databaseConnection.DB()
	if err != nil {
		return nil, err
	}

	// Connection pool settings from environment variables
	maxIdleConns, _ := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", "10"))
	maxOpenConns, _ := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", "25"))
	connMaxLifetime, _ := strconv.Atoi(getEnv("DB_CONN_MAX_LIFETIME", "3600")) // in seconds

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	return databaseConnection, nil
}


// performMigrations performs the database migrations
func performMigrations() error {
	err := dataBaseInstance.AutoMigrate(&models.Countries{})
	return err
}

// getEnv retrieves the value of an environment variable, falling back to a default value
func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}


// While GORM does handle connection pooling automatically, explicitly closing the connection pool can be a good practice for clarity and resource management. It ensures that all resources are properly released when your application is shutting down.
// CloseDB closes the database connection pool
func CloseDB() {
	sqlDB, err := dataBaseInstance.DB()
	if err != nil {
		log.Printf("Error retrieving database instance: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing database connection pool: %v", err)
	}
}


// Non-Destructive Column Operations (These operations modify the schema without affecting the existing data in the table.)
// 1. Adding New Columns.
// 2. Changing Column Attributes i.e not null, autoIncrement, etc.
// 3. Adding Indexes

// Destructive Column Operations
// 1. Dropping Columns (column dropped will lose its data)
// 2. Renaming Columns
// 3. Changing Column Data Types

