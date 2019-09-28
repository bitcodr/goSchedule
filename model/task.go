package model

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
	ExecuteDate uint32
	Reference   string
	Email       Email
}
