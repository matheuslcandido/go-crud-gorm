package database

import (
	"go-crud-gorm/entity"
	"go-crud-gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func Connect() error {
	logger.Info.Print("Starting connect on DB")

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error.Print(err)
		return err
	}

	Db = db

	logger.Info.Print("DB connected")

	logger.Info.Print("Starting execute migrations")

	err = Db.AutoMigrate(&entity.Product{})
	if err != nil {
		logger.Error.Print(err)
		return err
	}

	logger.Info.Print("Migrations executed")

	return nil
}
