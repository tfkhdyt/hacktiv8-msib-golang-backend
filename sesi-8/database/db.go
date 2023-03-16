package database

import (
	"fmt"
	"log"
	"sesi-8/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

const (
	host     = "localhost"
	user     = "tfkhdyt"
	password = "69420"
	dialect  = "postgres"
	port     = "5432"
	dbname   = "h8-movies"
)

func handleDatabaseConnection() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err = gorm.Open(postgres.Open(psqlInfo))
	if err != nil {
		log.Panicln(err.Error())
	}
}

func migrateModel() {
	if err := db.AutoMigrate(&entity.Movie{}); err != nil {
		log.Panicln(err.Error())
	}
}

func InitializeDatabase() {
	handleDatabaseConnection()
	migrateModel()
}

func GetDatabaseInstance() *gorm.DB {
	return db
}
