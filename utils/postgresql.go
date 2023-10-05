package utils

import (
	"eticketing/entity"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	// connect using gorm pgx
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        os.Getenv("DATABASE_URL"),
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	_ = conn.AutoMigrate(
		entity.User{},
		entity.Terminal{},
	)

	SetupDBConnection(conn)
	return conn, nil
}

func SetupDBConnection(DB *gorm.DB) {
	db = DB
}

func GetDBConnection() *gorm.DB {
	return db
}
