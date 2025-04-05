package config

import (
	"os"

	"github.com/luispolippo/goportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/openings.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, Creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("create directory error: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("create database file error: %v", err)
			return nil, err
		}
		file.Close()
	}

	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("automigrate error: %v", err)
		return nil, err
	}

	return db, nil
}
