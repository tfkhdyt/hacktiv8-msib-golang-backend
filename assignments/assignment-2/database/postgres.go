package database

import (
	"log"

	"assignment_2/config/database_config"
	"assignment_2/entity"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open(database_config.GetDBConfig())
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err := db.AutoMigrate(&entity.Order{}, &entity.Item{}); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Connected to DB!")
}

func GetPostgresInstance() *gorm.DB {
	return db
}
