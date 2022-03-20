package executor

import (
	"log"
	"sync/atomic"
)

// 任务执行器
// maxWorkers: 同时执行的goroutine最大数量
// 'ActiveWorkers' tells the number of active goroutines spawned by the Executor at given time.
// Tasks: 执行器接收任务的channel
// Reports: 执行器发布报告的channel
// signals: 可以控制执行器的channel，目前仅支持发送终止信号(1)
type Executor struct {
	maxWorkers    int64
	ActiveWorkers int64

	Tasks   chan Task
	Reports chan Report
	signals chan int
}

// 创建一个任务执行器
func NewExecutor(maxWorkers int, signals chan int) *Executor {
	chanSize := 100

	if maxWorkers > chanSize {
		chanSize = maxWorkers
	}

	executor := Executor{
		maxWorkers: int64(maxWorkers),
		Tasks:      make(chan Task, chanSize),
		Reports:    make(chan Report, chanSize),
		signals:    signals,
	}

	go executor.launch()

	return &executor
}

// 启动主循环， 拉取相关channel并处理不同message
func (executor *Executor) launch() int {
	reports := make(chan Report, executor.maxWorkers)

	for {
		select {
		case signal := <-executor.signals:
			if executor.handleSignals(signal) == 0 {
				return 0
			}

		case r := <-reports:
			executor.addReport(r)

		default:
			if executor.ActiveWorkers < executor.maxWorkers && len(executor.Tasks) > 0 {
				task := <-executor.Tasks
				atomic.AddInt64(&executor.ActiveWorkers, 1)
				go executor.launchWorker(task, reports)
			}
		}
	}
}

// handleSignals is called whenever anything is received on the 'signals' channel.
// It performs the relevant task according to the received signal(request) and then responds either
// with 0 or 1 indicating whether the request was respected(0) or rejected(1).
func (executor *Executor) handleSignals(signal int) int {
	if signal == 1 {
		log.Println("Received termination request...")

		if executor.Inactive() {
			log.Println("No active workers, exiting...")
			executor.signals <- 0
			return 0
		}

		executor.signals <- 1
		log.Println("Some tasks are still active...")
	}

	return 1
}

// 收到task并启动新worker来执行，将report发布到内部channel
func (executor *Executor) launchWorker(task Task, reports chan<- Report) {
	report := task.Execute()

	if len(reports) < cap(reports) {
		reports <- report
	} else {
		log.Println("Executor's report channel is full...")
	}

	atomic.AddInt64(&executor.ActiveWorkers, -1)
}

// 使用非阻塞方式将task提交到Executor任务执行器. 
// 如果任务执行器中的task channel已满，client应考虑重试提交。
func (executor *Executor) AddTask(task Task) bool {
	if len(executor.Tasks) == cap(executor.Tasks) {
		return false
	}

	executor.Tasks <- task
	return true
}

// addReport is used by the Executor to publish the reports in a non-blocking way. It client is not
// reading the reports channel or is slower that the Executor publishing the reports, the Executor's
// reports channel is going to get full. In that case this method will not block and that report will
// not be added.
func (executor *Executor) addReport(report Report) bool {
	if len(executor.Reports) == cap(executor.Reports) {
		return false
	}

	executor.Reports <- report
	return true
}

// Inactive checks if the Executor is idle. This happens when there are no pending tasks, active
// workers and reports to publish.
func (executor *Executor) Inactive() bool {
	return executor.ActiveWorkers == 0 && len(executor.Tasks) == 0 && len(executor.Reports) == 0
}