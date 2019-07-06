package provider

import "github.com/amiraliio/goSchedule/handler"

//App set application service handlers
type App struct {
	Task handler.TaskService
}
