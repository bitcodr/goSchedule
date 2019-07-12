package interfaces

import (
	"context"

	"github.com/amiraliio/goSchedule/model"
)

type baseRepositoryInterFace interface {
	List(context context.Context, filter *model.Filter) []*model.Task
	Get(context context.Context, filter *model.Filter, id string) []*model.Task
}
