package helpers

import (
	"context"
	"time"
)

//ContextTimeout helper
func ContextTimeout(timeout int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
}