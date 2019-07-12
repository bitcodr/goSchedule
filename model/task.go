package model

//TaskCollection is task collection name costant
const TaskCollection string = "test"

//Pending is status of pending task
const Pending string = "p"

//Done is status of done task
const Done string = "d"

//Task model
type Task struct {
	Executedate int
	Reference   string
	Email       Email
}
