package models

//Email model
type Email struct {
	Attachments []string
	Body        string
	Receiver    string
	Subject     string
}

//Task model
type Task struct {
	Executedate int
	Reference   string
	Email       Email
}
