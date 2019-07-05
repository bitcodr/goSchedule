package config

import (
	"log"

	"github.com/amiraliio/goSchedule/handler"
	"github.com/joho/godotenv"
)

//AppProvider set application service handlers
type AppProvider struct {
	Task handler.TaskService
}


//Instantiate application configs
func Instantiate() {
	env()
}

func env() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file doesn't loaded")
	}
}
