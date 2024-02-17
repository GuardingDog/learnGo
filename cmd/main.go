package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentCtx := context.Background()

	// Using WithTimeout
	timeoutCtx, cancel := context.WithTimeout(parentCtx, 2*time.Second)
	defer cancel()

	// Using WithDeadline
	deadlineCtx, cancel := context.WithDeadline(parentCtx, time.Now().Add(2*time.Second))
	defer cancel()

	// Simulating long-running operations
	go func() {
		select {
		case <-timeoutCtx.Done():
			fmt.Println("Timeout Context canceled")
		}
	}()

	go func() {
		select {
		case <-deadlineCtx.Done():
			fmt.Println("Deadline Context canceled")
		}
	}()

	// Waiting for a while to observe the difference
	time.Sleep(3 * time.Second)
}
