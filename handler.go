package main

import (
	"github.com/amiraliio/goSchedule/helpers"
	"github.com/amiraliio/goSchedule/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const collection string = "test"
const records int = 3
const pending string = "p"
const done string = "d"

func schedule() {
	getTasks()
}

func getTasks() {
	collection := db().Collection(collection)
	ctx, _ := helpers.ContextTimeout(30)
	filter := bson.D{
		primitive.E{
			Key:   "status",
			Value: pending,
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
	findOptions.SetLimit(int64(records)).SetSort(primitive.M{
		"execute_date": 1,
	})
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer cursor.Close(ctx)
	var results []*models.Task
	for cursor.Next(ctx) {
		var task models.Task
		err := cursor.Decode(&task)
		results = append(results, &task)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	for _, v := range results {
		// if i == records - 1{
		//             fmt.Println(i)
		//     getTasks()
		//     return
		// }
		helpers.SendEmail(v.Email)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err.Error())
	}
}
