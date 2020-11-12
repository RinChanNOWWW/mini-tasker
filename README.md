# mini-tasker
一个简单的 Go 计时任务 Library。

主要用来给自己练手，缝合一些平时学到的功能实现。

## Usage

### pre

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


### 向调度器中加入任务

```go
task := MyTask{}
scheduler.AddTask(task)
```

### 启动调度器

```go
scheduler.Start()
// 之后加入的任务只能单独启动，或者停止全部任务重新启动
```

### 停止调度器

```go
scheduler.Stop()
```

### 启动单个任务

```go
scheduler.StartTask("MyTask")
```

### 停止单个任务

```go
scheduler.StopTask("MyTask")
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

## TODO
- [ ] 优化代码逻辑
- [ ] 测试更多 case，查找并修复 bug
- [ ] 考虑加入中间件功能


