package posgresql

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresGormDB(login, pass, host, port, dbName string) (*gorm.DB, error) {
	dbString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Europe/Kiev",
		login, pass, host, port, dbName,
	)

	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
