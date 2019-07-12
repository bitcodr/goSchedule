package model

//Email model
type Email struct {
	Attachments []string
	Body        string
	Receiver    string
	Subject     string
}
