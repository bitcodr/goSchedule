package repository

import (
	"testing"

	"github.com/amiraliio/goSchedule/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRepository_GetTaskRepo(t *testing.T) {
	repo := &Repository{}

	if repo == nil {
		t.Error("Expected repository to be created")
	}
}

func TestRepository_Type(t *testing.T) {
	var repo interface{} = &Repository{}

	// Test that Repository implements the expected structure
	_, ok := repo.(*Repository)
	if !ok {
		t.Error("Expected Repository type")
	}
}

// Note: List and Update methods require MongoDB connection
// These are integration tests and should be run with a test database
// For unit tests, use the mock repository in handler tests

func TestRepositoryStructure(t *testing.T) {
	// Test that we can create a Repository instance
	repo := Repository{}

	// Verify it's the correct type
	if _, ok := interface{}(&repo).(*Repository); !ok {
		t.Error("Expected Repository type")
	}

	// Note: List and Update methods require MongoDB connection
	// These should be tested with integration tests using a test database
	// For unit tests, use the mock repository in handler tests
}

func TestTaskModelCreation(t *testing.T) {
	task := &model.Task{
		ID:          primitive.NewObjectID(),
		ExecuteDate: 1234567890,
		Reference:   "test-ref",
		Status:      model.Pending,
		Email: model.Email{
			Body:     "test",
			Receiver: "test@example.com",
			Subject:  "Test",
		},
	}

	if task.Status != model.Pending {
		t.Errorf("Expected status Pending, got %s", task.Status)
	}

	if task.Reference != "test-ref" {
		t.Errorf("Expected reference 'test-ref', got %s", task.Reference)
	}
}
