package model

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaskCollection(t *testing.T) {
	if TaskCollection != "test" {
		t.Errorf("Expected TaskCollection to be 'test', got '%s'", TaskCollection)
	}
}

func TestStatusConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant string
		expected string
	}{
		{"Pending status", Pending, "p"},
		{"Done status", Done, "d"},
		{"Failed status", Failed, "f"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, tt.constant)
			}
		})
	}
}

func TestTaskModel(t *testing.T) {
	email := Email{
		Attachments: []string{"file1.pdf", "file2.pdf"},
		Body:        "Test email body",
		Receiver:    "test@example.com",
		Subject:     "Test Subject",
	}

	task := Task{
		ID:          primitive.NewObjectID(),
		ExecuteDate: 1234567890,
		Reference:   "test-ref-123",
		Status:      Pending,
		Email:       email,
	}

	if task.Status != Pending {
		t.Errorf("Expected status to be Pending, got %s", task.Status)
	}

	if task.Reference != "test-ref-123" {
		t.Errorf("Expected reference to be 'test-ref-123', got %s", task.Reference)
	}

	if task.ExecuteDate != 1234567890 {
		t.Errorf("Expected executeDate to be 1234567890, got %d", task.ExecuteDate)
	}

	if task.Email.Receiver != "test@example.com" {
		t.Errorf("Expected email receiver to be 'test@example.com', got %s", task.Email.Receiver)
	}
}
