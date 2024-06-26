package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize MySQL
	db, err = InitializeMySQL()
	if err != nil {
		return fmt.Errorf("error initializing MySQL: %v", err)
	}

	return nil
}

func GetMySQL() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
