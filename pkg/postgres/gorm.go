package postgres

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// Once method is used to prevent race condition between go routines in case the lazy db initializer is called inside some of them
	once sync.Once

	instance      *gorm.DB
	connectionStr string
)

// SetConnectionString sets the connection string for the database connection.
func SetConnectionString(dsn string) {
	connectionStr = dsn
}

// GetInstance retrieves the current GORM database instance.
func GetInstance() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		err = connect()
	})
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// Migrate runs the migrations.
func Migrate(models ...interface{}) error {
	if err := connect(); err != nil {
		return fmt.Errorf("error connecting to database for migration: %w", err)
	}

	if err := instance.AutoMigrate(models...); err != nil {
		return fmt.Errorf("error migrating database: %w", err)
	}

	return nil
}

// Close gracefully closes the database connection.
func Close() {
	if err := closeDB(instance); err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}
}

// connect establishes a connection to the database if it isn't already connected.
func connect() error {
	if instance != nil {
		return nil
	}

	if connectionStr == "" {
		return fmt.Errorf("connection string is not set")
	}

	var err error
	instance, err = open(connectionStr)
	if err != nil {
		return fmt.Errorf("error getting DB instance: %w", err)
	}
	return nil
}

// open creates a new GORM database connection using the given DSN.
func open(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// closeDB handles closing the GORM database instance.
func closeDB(gormDB *gorm.DB) error {
	if gormDB == nil {
		return nil
	}

	db, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("error getting DB instance: %w", err)
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf("error closing DB instance: %w", err)
	}

	return nil
}
