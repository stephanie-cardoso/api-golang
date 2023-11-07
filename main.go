package main

import (
	"github.com/rs/zerolog/log"
	"github.com/stephanie-cardoso/api-golang/config"
	"github.com/stephanie-cardoso/api-golang/router"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to get api configs")
		return
	}
	router.Initialize()
}
