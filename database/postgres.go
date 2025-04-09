package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost port=5432 user=postgres password=123@senha dbname=teste-go sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}

// GetSQLDB returns the underlying *sql.DB from the GORM database connection
func GetSQLDB(db *gorm.DB) (*gorm.DB, error) {
	return db, nil
}
