package handler

import (
	"context"

	"github.com/amiraliio/goSchedule/helper"
	"github.com/amiraliio/goSchedule/model"
	"github.com/amiraliio/goSchedule/repository"
)

//TaskService to use other repository or database session
type TaskService struct {
	ctx context.Context
}

func (s *TaskService) getTaskRepo() repository.TaskRepo {
	return &repository.Repository{}
}

//List of tasks handler
func (s *TaskService) List() {
	filter := &model.Filter{
		Limit: 3,
	}
	results := s.getTaskRepo().Get(ctx, filter)

	for _, v := range results {
		// if i == records - 1{
		//             fmt.Println(i)
		//     getTasks()
		//     return
		// }
		helper.SendEmail(v.Email)
	}

}
