package repository

import (
	"context"
	"log"
	"time"

	"github.com/amiraliio/goSchedule/config"
	"github.com/amiraliio/goSchedule/helper"
	"github.com/amiraliio/goSchedule/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Repository struct for sessions like database session
type Repository struct{}

//List of tasks from mongo
func (r *Repository) List(context context.Context, filter *model.Filter) []*model.Task {
	collection := config.DB().Collection(model.TaskCollection)
	ctx, _ := helper.ContextTimeout(30)
	query := bson.D{
		primitive.E{
			Key:   "status",
			Value: model.Pending,
		},
		primitive.E{
			Key: "executeDate",
			Value: bson.D{
				primitive.E{
					Key:   "$lte",
					Value: time.Now().Unix(),
				},
			},
		},
	}
	findOptions := options.Find()
	findOptions.SetLimit(filter.Limit).SetSort(primitive.M{
		"executeDate": 1,
	})
	cursor, err := collection.Find(ctx, query, findOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer cursor.Close(ctx)
	var results []*model.Task
	for cursor.Next(ctx) {
		var task model.Task
		err := cursor.Decode(&task)
		results = append(results, &task)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err.Error())
	}
	return results
}

//Update from mongo database
func (r *Repository) Update(context context.Context, task *model.Task) *model.Task {
	collection := config.DB().Collection(model.TaskCollection)
	ctx, _ := helper.ContextTimeout(30)
	id := bson.D{
		primitive.E{
			Key:   "_id",
			Value: task.ID,
		},
	}
	model := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{
					Key:   "status",
					Value: model.Done,
				},
			},
		},
	}
	updated, err := collection.UpdateOne(ctx, id, model)
	if err != nil {
		log.Println("error in updated")
	}
	if updated.UpsertedCount > 0 {
		return task
	}
	return nil
}
