package config

import (
	"log"

	"github.com/joho/godotenv"
)

//Instantiate application configs
func Instantiate() {
	env()
}

func env() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file couldn't loaded")
	}
}
