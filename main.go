package main

import (
	"github.com/amiraliio/goSchedule/config"
	"github.com/amiraliio/goSchedule/provider"
)

func main() {
	config.Instantiate()
	app := new(provider.App)
	app.Task.List()
}
