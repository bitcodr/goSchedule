package provider

import (
	"github.com/amiraliio/goSchedule/handler"
)

//providers
var (
	TaskService *handler.TaskService
)

func init() {
	TaskService = handler.NewTaskService()
}
