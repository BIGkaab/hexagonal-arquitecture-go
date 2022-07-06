package config

import (
	"sync"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *gorm.DB
var once sync.Once

func ConnInstance() *gorm.DB {
	once.Do(func() {
		instance = getConnection()
	})

	return instance
}

func getConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=123456 dbname=test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	log.Info("Database successfully connected")

	return db

}
