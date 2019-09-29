package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TaskCollection is task collection name costant
const TaskCollection string = "test"

//Pending is status of pending task
const Pending string = "p"

//Done is status of done task
const Done string = "d"

//Failed is status of failed task
const Failed string = "f"

//Task model
type Task struct {
	ID          primitive.ObjectID
	ExecuteDate int64
	Reference   string
	Status      string
	Email       Email
}
