package db

import (
	"fmt"
	"log"
	"time"

	"github.com/twonkista/PikePlace/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open(dsn string) (*gorm.DB, error) {
	// Load environment variables
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL is empty")
	}

	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // dev-friendly logging
	})

	err = gdb.AutoMigrate(&models.User{}, &models.Pool{}, &models.Wager{})
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	if err != nil {
		return nil, fmt.Errorf("gorm open: %w", err)
	}

	// Optional but useful: tune and ping the underlying pool
	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, fmt.Errorf("get underlying sql db: %w", err)
	}

	// Reasonable defaults (adjust later)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return gdb, nil
}
