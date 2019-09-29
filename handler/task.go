package handler

import (
	"context"
	"fmt"

	"github.com/amiraliio/goSchedule/helper"
	"github.com/amiraliio/goSchedule/model"
	"github.com/amiraliio/goSchedule/repository"
	"github.com/amiraliio/goSchedule/repository/interfaces"
)

//TaskService to use other repository or database session
type TaskService struct {
	ctx context.Context
}

func (s *TaskService) getTaskRepo() interfaces.TaskRepo {
	return &repository.Repository{}
}

//List of tasks handler
func (s *TaskService) List() {
	filter := &model.Filter{
		Limit: 3,
	}
	results := s.getTaskRepo().List(s.ctx, filter)

	for _, v := range results {
		if helper.SendEmail(v.Email) {
			fmt.Println(true)
			continue
			//TODO update the record state to active in the database collection
		}
		fmt.Println(false)
		continue
		//TODO update the record state to failed in the database collection
	}

}