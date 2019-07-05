package main

import (
	"github.com/amiraliio/goSchedule/config"
)

func main() {
	config.Instantiate()
	app := new(config.AppProvider)
	app.Task.List()
}
