package models

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbhost = "localhost"
	dbuser = "postgres"
	dbpass = "postgres"
	dbport = "5432"
	dbname = "ordersitems"
	db     *gorm.DB
	err    error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbhost, dbuser, dbpass, dbname, dbport)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.Debug().AutoMigrate(&Order{}, &Item{})
}

func GetDB() *gorm.DB {
	return db
}
