package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initConnection() {
	dsn := "host=localhost user=user dbname=db password=password sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
