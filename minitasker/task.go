package minitasker

import "time"

// Task is the base interface of the task
// the task should impl these methods
type Task interface {
	// DoTask is what you want the task to do
	DoTask()
	// Interval returns the frequency of doing the task
	Interval() time.Duration
	// TaskName returns the name of your task
	TaskName() string
}
