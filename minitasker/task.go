package minitasker

import "time"

// Task is the base interface of the task
// the task should impl these methods
type Task interface {
	DoTask()
	Interval() time.Duration
	TaskName() string
}
