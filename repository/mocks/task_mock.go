package mocks

import (
	"context"

	"github.com/amiraliio/goSchedule/model"
)

// MockTaskRepository is a mock implementation of TaskRepo interface for testing
type MockTaskRepository struct {
	ListFunc   func(ctx context.Context, filter *model.Filter) []*model.Task
	UpdateFunc func(ctx context.Context, task *model.Task) *model.Task
}

// List calls the mock ListFunc
func (m *MockTaskRepository) List(ctx context.Context, filter *model.Filter) []*model.Task {
	if m.ListFunc != nil {
		return m.ListFunc(ctx, filter)
	}
	return []*model.Task{}
}

// Update calls the mock UpdateFunc
func (m *MockTaskRepository) Update(ctx context.Context, task *model.Task) *model.Task {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, task)
	}
	return task
}
