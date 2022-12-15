package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	setupDatabaseConnection("db/data.db")
}

func ConnectDatabaseTesting() {
	setupDatabaseConnection("file::memory:?cache=shared")
}

func setupDatabaseConnection(databaseLocation string){
	database, err := gorm.Open(sqlite.Open(databaseLocation), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&User{})
	if err != nil {
		return
	}

	err = database.AutoMigrate(&Book{})
	if err != nil {
		return
	}

	DB = database
}