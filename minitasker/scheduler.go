package minitasker

import (
	"context"
	"sync"
	"time"
)

// Scheduler is the struct that holds all tasks
type Scheduler struct {
	tasks      map[string]Task
	contextMap map[string]context.Context
	cancelMap  map[string]context.CancelFunc
	runMap     map[string]bool
	ctx        context.Context

	sync.Mutex
}

// NewScheduler returns an instantiated Scheduler
func NewScheduler(ctx context.Context) *Scheduler {
	s := &Scheduler{}
	s.tasks = make(map[string]Task)
	s.contextMap = make(map[string]context.Context)
	s.cancelMap = make(map[string]context.CancelFunc)
	s.runMap = make(map[string]bool)
	s.ctx = ctx
	return s
}

// AddTask add a task into the Scheduler
func (s *Scheduler) AddTask(task Task) error {
	s.Lock()
	defer s.Unlock()
	s.tasks[task.TaskName()] = task
	taskName := task.TaskName()
	if _, ok := s.runMap[taskName]; ok {
		return taskError(errAddDupNameTask, taskName)
	}
	s.runMap[taskName] = false
	return nil
}

// Start is the method to start all added Task
func (s *Scheduler) Start() {
	for _, task := range s.tasks {
		// the task is not running
		if !s.runMap[task.TaskName()] {
			childCtx, cancel := context.WithCancel(s.ctx)
			s.contextMap[task.TaskName()] = childCtx
			s.cancelMap[task.TaskName()] = cancel
			go s.loop(childCtx, task)
		}
	}
}

// loop is the method for task to run periodically
func (s *Scheduler) loop(ctx context.Context, task Task) {
	s.Lock()
	s.runMap[task.TaskName()] = true
	s.Unlock()
	ticker := time.NewTicker(task.Interval())
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			s.Lock()
			s.runMap[task.TaskName()] = false
			s.Unlock()
			return
		case <-ticker.C:
			task.DoTask()
		}
	}
}

// Stop is the method to stop all the task
func (s *Scheduler) Stop() {
	for _, task := range s.tasks {
		if s.runMap[task.TaskName()] {
			s.cancelMap[task.TaskName()]()
		}
	}
}

// StartTask is the method to start a determined task
func (s *Scheduler) StartTask(taskName string) error {
	if run, ok := s.runMap[taskName]; ok {
		if !run {
			childCtx, cancel := context.WithCancel(s.ctx)
			s.contextMap[taskName] = childCtx
			s.cancelMap[taskName] = cancel
			go s.loop(childCtx, s.tasks[taskName])
		}
		return nil
	}
	return taskError(errNoTask, taskName)
}

// StopTask is the method to stop a determined task
func (s *Scheduler) StopTask(taskName string) error {
	if cancel, ok := s.cancelMap[taskName]; ok {
		cancel()
		return nil
	}
	return taskError(errNoTask, taskName)
}
