package main

import (
	"fmt"
	"github.com/amiraliio/goSchedule/config"
	"github.com/amiraliio/goSchedule/provider"
)

func main() {
	config.Env()
	provider.TaskService.List()
	fmt.Println("Done")
}
