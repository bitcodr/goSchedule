package config

import (
	"log"

	"github.com/joho/godotenv"
)

//Env config function
func Env() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file couldn't loaded")
	}
}
