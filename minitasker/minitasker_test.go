package minitasker

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type MyTask1 struct {
	Count int
}

func (t MyTask1) DoTask() {
	fmt.Println("Do Task1", time.Now().Format("2006-01-02 15:04:05"))
}

func (t MyTask1) Interval() time.Duration {
	return time.Second
}

func (t MyTask1) TaskName() string {
	return "MyTask1"
}

type MyTask2 struct {
}

func (t MyTask2) DoTask() {
	fmt.Println("Do Task2", time.Now().Format("15:04:05"))
}

func (t MyTask2) Interval() time.Duration {
	return time.Second
}

func (t MyTask2) TaskName() string {
	return "MyTask2"
}

func TestStartAndStopTask(t *testing.T) {
	s := NewScheduler(context.Background())
	task1 := MyTask1{Count: 0}
	s.AddTask(task1)
	fmt.Println("Start Task1")
	s.StartTask(task1.TaskName())
	time.Sleep(5 * time.Second)
	s.Stop()
	fmt.Println("Stop Task1")
	time.Sleep(5 * time.Second)
	fmt.Println("Start Task1")
	s.StartTask(task1.TaskName())
	time.Sleep(5 * time.Second)
	s.Stop()
}
