package config

import (
	"os"

	"github.com/tmazleo/opportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	//create database and connect
	dbPath := "./db/main.db"
	//checking if db file exist
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database is not exist. Creating...")
		//creating db path in permanent mode
		err = os.MkdirAll("./db", os.ModePerm)
		//return error
		if err != nil {
			return nil, err
		}
		//crating file with path from dbPath
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		//closing file !essential!
		file.Close()
	}

	//qual database esta utilizando
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.ErrorF("sqlite openning error: %v", err)
		return nil, err
	}
	//migrate schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrorF("sqlite automigration error: %v", err)
		return nil, err
	}
	//return db
	return db, nil
}