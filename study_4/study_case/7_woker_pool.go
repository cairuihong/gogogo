package study_case

import (
	"fmt"
	"math/rand"
)

/*
worker pool（goroutine池）----线程池
	​​资源复用​​：重用已创建的 goroutine，减少创建销毁开销。
	​​并发控制​​：通过固定池大小限制同时运行的 goroutine 数量，避免系统过载。
	​​任务队列管理​​：缓冲未处理任务，平滑流量峰值。
*/


// Job：表示待处理任务，包含ID和随机数。
type Job struct {
	Id      int // 任务唯一标识符
	RandNum int // 随机数
}

// Result：表示任务处理结果，包含原任务和数字各位数之和。
type Result struct {
	job *Job // 指向原始任务的指针
	sum int  // 随机数各位数字之和
}


// WorderPoolCase：主函数，不断生成随机数任务放入任务通道，并启动goroutine打印结果通道中的处理结果。
// 该函数演示了goroutine池的工作流程，包括任务分发、并发处理和结果收集。
func WorderPoolCase() {
	jobChan := make(chan *Job, 128)     // 创建带缓冲的任务通道，用于存放待处理任务
	resultChan := make(chan *Result, 128) // 创建带缓冲的结果通道，用于存放处理结果

	createPool(64, jobChan, resultChan) // 创建包含64个goroutine的工作池

	// 启动一个goroutine来消费结果通道中的数据并打印
	go func(ResultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id %d ,randnum %d result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)

	var id int
	// 不断生成新的任务并发送到任务通道
	for {
		id++
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}

// createPool：创建指定数量的goroutine作为工作池，从任务通道读取任务并计算结果发送到结果通道。
// num: 要创建的goroutine数量
// jobChan: 用于接收待处理任务的通道
// resultChan: 用于发送处理结果的通道
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 每个goroutine持续从任务通道中读取任务
			for job := range jobChan {
				r_num := job.RandNum

				// 计算随机数各位数字之和
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}

				// 构造结果并发送到结果通道
				r := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}