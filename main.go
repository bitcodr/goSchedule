package main

import (
	"fmt"
	"github.com/amiraliio/goSchedule/config"
	"github.com/amiraliio/goSchedule/provider"
)

func main() {
	config.Instantiate()
	app := new(provider.App)
	app.Task.List()
	fmt.Println("Done")
}
