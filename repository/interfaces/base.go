package interfaces

import (
	"context"

	"github.com/amiraliio/goSchedule/model"
)

type baseRepositoryInterFace interface {
	List(context context.Context, filter *model.Filter) []*model.Task
	Update(context context.Context, task *model.Task) *model.Task
}
