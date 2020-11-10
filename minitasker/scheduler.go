package minitasker

import (
	"context"
	"errors"
	"time"
)

var (
	// ErrAddDupNameTask informs their is already a task in the
	// scheduler with the same name.
	ErrAddDupNameTask = errors.New("Duplicated Task Name")
)

// scheduler is the struct that holds all tasks
type scheduler struct {
	tasks      []Task
	contextMap map[string]context.Context
	cancelMap  map[string]context.CancelFunc
	ctx        context.Context
}

// NewScheduler returns an instantiated scheduler
func NewScheduler(ctx context.Context) *scheduler {
	s := &scheduler{}
	s.tasks = make([]Task, 0)
	s.contextMap = make(map[string]context.Context)
	s.cancelMap = make(map[string]context.CancelFunc)
	s.ctx = ctx
	return s
}

// AddTask add a task into the scheduler
func (s *scheduler) AddTask(task Task) error {
	s.tasks = append(s.tasks, task)
	taskName := task.TaskName()
	if _, ok := s.cancelMap[taskName]; ok {
		return ErrAddDupNameTask
	}
	childCtx, cancel := context.WithCancel(s.ctx)
	s.contextMap[taskName] = childCtx
	s.cancelMap[taskName] = cancel
	return nil
}

// Start is the method to start all added Task
func (s *scheduler) Start() {
	for _, task := range s.tasks {
		ctx := s.contextMap[task.TaskName()]
		go s.loop(ctx, task)
	}
}

// loop is the method for task to run periodically
func (s *scheduler) loop(ctx context.Context, task Task) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(task.Interval())
			task.DoTask()
		}
	}
}

// Stop is the method to stop all the task
func (s *scheduler) Stop() {
	for _, task := range s.tasks {
		s.cancelMap[task.TaskName()]()
	}
}
