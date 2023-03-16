package database

import (
	"assignment_2/config/database_config"
	"log"

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
}

func GetPostgresInstance() *gorm.DB {
	return db
}
