package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"../minitasker"
)

type demoTask1 struct {
}

func (t *demoTask1) DoTask() {
	fmt.Println("task 1", time.Now().Format("2006-01-02 15:04:05"))
}

func (t *demoTask1) Interval() time.Duration {
	return time.Second
}

func (t *demoTask1) TaskName() string {
	return "task1"
}

type demoTask2 struct {
}

func (t *demoTask2) DoTask() {
	fmt.Println("task 2", time.Now().Format("2006-01-02 15:04:05"))
}

func (t *demoTask2) Interval() time.Duration {
	return time.Second * 2
}

func (t *demoTask2) TaskName() string {
	return "task2"
}

func main() {
	t1 := demoTask1{}
	t2 := demoTask2{}
	scheduler := minitasker.NewScheduler(context.Background())
	err := scheduler.AddTask(&t1)
	if err != nil {
		fmt.Println(err)
	}
	err = scheduler.AddTask(&t2)
	if err != nil {
		fmt.Println(err)
	}
	scheduler.Start()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Second)
		scheduler.Stop()
		fmt.Println("scheduler stopped.")
		time.Sleep(5 * time.Second)
	}()
	wg.Wait()
	fmt.Println("done.", time.Now().Format("2006-01-02 15:04:05"))
}
