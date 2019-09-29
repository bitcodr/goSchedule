package interfaces

import (
	"context"

	"github.com/amiraliio/goSchedule/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type baseRepositoryInterFace interface {
	List(context context.Context, filter *model.Filter) []*model.Task
	Update(context context.Context, ID primitive.ObjectID) []*model.Task
}
