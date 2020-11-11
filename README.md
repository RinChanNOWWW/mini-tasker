# mini-tasker
一个简单的 Go 计时任务包

## Usage

1. 定义自己的任务结构体，并实现 `Task` 接口

```go
// Example
type MyTask struct {
	// members you want...
}

// DoTask is what you want the task to do
func (t MyTask) DoTask {
  fmt.Println("Do my Task", time.Now().Format("2006-01-02 15:04:05"))
}

// Interval returns the frequency of doing the task
func (t MyTask) Interval() time.Duration {
  return time.Second
}

// TaskName returns the name of your task
func (t MyTask) TaskName() string {
  return "MyTask"
}
```

2. 实例化调度器

```go
scheduler := minitasker.NewScheduler(context.Background())
```

3. 向调度器中加入任务

```go
task := MyTask{}
scheduler.AddTask(task)
```

4. 启动调度器

```go
scheduler.Start()
```

5. 停止调度器

```go
scheduler.Stop()
```

## Run

```
Do my Task 2020-11-11 10:47:23
Do my Task 2020-11-11 10:47:24
Do my Task 2020-11-11 10:47:25
Do my Task 2020-11-11 10:47:26
Do my Task 2020-11-11 10:47:27
Do my Task 2020-11-11 10:47:28
...
```

