package handler

import (
	"context"
	"testing"

	"github.com/amiraliio/goSchedule/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewTaskService(t *testing.T) {
	service := NewTaskService()

	if service == nil {
		t.Error("Expected TaskService to be created")
	}

	if service.ctx == nil {
		t.Error("Expected context to be initialized")
	}
}

func TestTaskService_GetTaskRepo(t *testing.T) {
	service := NewTaskService()
	repo := service.getTaskRepo()

	if repo == nil {
		t.Error("Expected repository to be created")
	}
}

func TestTaskServiceStructure(t *testing.T) {
	service := &TaskService{
		ctx: context.Background(),
	}

	if service.ctx == nil {
		t.Error("Expected context to be set")
	}

	// Test that getTaskRepo returns a repository
	repo := service.getTaskRepo()
	if repo == nil {
		t.Error("Expected repository to not be nil")
	}
}

// MockEmail function for testing without actually sending emails
var testEmailSent bool

func TestTaskCreation(t *testing.T) {
	task := &model.Task{
		ID:          primitive.NewObjectID(),
		ExecuteDate: 1234567890,
		Reference:   "test-ref-123",
		Status:      model.Pending,
		Email: model.Email{
			Body:     "Test email body",
			Receiver: "test@example.com",
			Subject:  "Test Subject",
		},
	}

	if task.Status != model.Pending {
		t.Errorf("Expected status Pending, got %s", task.Status)
	}

	// Update status to Done
	task.Status = model.Done
	if task.Status != model.Done {
		t.Errorf("Expected status Done, got %s", task.Status)
	}

	// Update status to Failed
	task.Status = model.Failed
	if task.Status != model.Failed {
		t.Errorf("Expected status Failed, got %s", task.Status)
	}
}

func TestTaskServiceContext(t *testing.T) {
	service := NewTaskService()

	// Verify context is background context
	select {
	case <-service.ctx.Done():
		t.Error("Context should not be done")
	default:
		// Context is not done, as expected
	}
}
