package helper

import (
	"context"
	"testing"
	"time"
)

func TestContextTimeout(t *testing.T) {
	tests := []struct {
		name    string
		timeout int
		want    time.Duration
	}{
		{
			name:    "30 seconds timeout",
			timeout: 30,
			want:    30 * time.Second,
		},
		{
			name:    "60 seconds timeout",
			timeout: 60,
			want:    60 * time.Second,
		},
		{
			name:    "1 second timeout",
			timeout: 1,
			want:    1 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := ContextTimeout(tt.timeout)
			defer cancel()

			deadline, ok := ctx.Deadline()
			if !ok {
				t.Error("Expected context to have a deadline")
				return
			}

			// Check if the deadline is approximately correct (within 100ms tolerance)
			expectedDeadline := time.Now().Add(tt.want)
			diff := expectedDeadline.Sub(deadline)
			if diff < 0 {
				diff = -diff
			}
			if diff > 100*time.Millisecond {
				t.Errorf("Deadline difference too large: %v", diff)
			}
		})
	}
}

func TestContextTimeoutCancellation(t *testing.T) {
	ctx, cancel := ContextTimeout(1)

	// Cancel immediately
	cancel()

	// Check if context is cancelled
	select {
	case <-ctx.Done():
		if ctx.Err() != context.Canceled {
			t.Errorf("Expected context.Canceled, got %v", ctx.Err())
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Context was not cancelled")
	}
}

func TestContextTimeoutExpiration(t *testing.T) {
	ctx, cancel := ContextTimeout(1)
	defer cancel()

	// Wait for context to timeout
	select {
	case <-ctx.Done():
		if ctx.Err() != context.DeadlineExceeded {
			t.Errorf("Expected context.DeadlineExceeded, got %v", ctx.Err())
		}
	case <-time.After(2 * time.Second):
		t.Error("Context did not timeout as expected")
	}
}
