package database

import (
	"gatorcan-backend/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect establishes a connection to the database
func Connect(config config.DatabaseConfig) (*gorm.DB, error) {
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	// 	config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := gorm.Open(sqlite.Open("file:e-learning.db"), &gorm.Config{}, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Migrate performs database schema migration
func Migrate(db *gorm.DB, models ...interface{}) error {
	return db.AutoMigrate(models...)
}
