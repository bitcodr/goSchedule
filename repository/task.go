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
func List(context context.Context, filter *model.Filter) []*model.Task {
	collection := config.DB().Collection(model.TaskCollection)
	ctx, _ := helper.ContextTimeout(30)
	query := bson.D{
		primitive.E{
			Key:   "status",
			Value: model.Pending,
		},
		primitive.E{
			Key: "execute_date",
			Value: bson.D{
				primitive.E{
					Key:   "$lt",
					Value: time.Now().Unix(),
				},
			},
		},
	}
	findOptions := options.Find()
	findOptions.SetLimit(int64(filter.Limit)).SetSort(primitive.M{
		"execute_date": 1,
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

//Get from mongo database
func Get(context context.Context, filter *model.Filter, id string) []*model.Task {
	return nil
}
