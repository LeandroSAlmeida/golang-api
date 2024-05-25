package config

import (
	"github.com/LeandroSAlmeida/golang-api/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMySQL() (*gorm.DB, error) {
	// Get logger instance
	logger := GetLogger("mysql")

	// Configure MySQL connection string
	dsn := "root:root@tcp(localhost:3306)/golang_api?charset=utf8mb4&parseTime=True&loc=Local"

	// Open connection to MySQL database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("mysql opening error: %v", err)
		return nil, err
	}

	// Additional check to ensure database is initialized
	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorf("error getting generic database object: %v", err)
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		logger.Errorf("error pinging database: %v", err)
		return nil, err
	}
	logger.Info("database connection verified successfully")

	// Migrate the schema
	logger.Info("running AutoMigrate")
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("mysql automigration error: %v", err)
		return nil, err
	}

	// Return the database connection
	return db, nil
}
