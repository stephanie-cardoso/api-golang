package config

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	db, err = InitializeSQLite()
	if err != nil {
		log.Error().Err(err).Msgf("[config] Failed to initialize SQLite")
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}
