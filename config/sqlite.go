package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/stephanie-cardoso/api-golang/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		log.Info().Msgf("[config] Database file not found, creating...")
		// Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msgf("[config] Failed to initialize SQLite")
		return nil, err
	}

	if err = db.AutoMigrate(&schemas.Opening{}); err != nil {
		log.Error().Err(err).Msgf("[config] Failed to auto migrate SQLite")
		return nil, err
	}

	return db, nil
}
