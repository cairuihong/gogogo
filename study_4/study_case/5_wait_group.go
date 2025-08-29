package study_case

/*

sync.WaitGroup 用于等待多个 Goroutine 完成
核心思想是通过计数器机制实现并发任务的等待
Add/Done/Wait 三部曲

*/

import (
	"fmt"
	"sync"
	"time"
)

// worker 是一个工作协程函数，用于执行具体的任务
// id: 工作协程的标识符
// wg: WaitGroup指针，用于同步协程完成状态
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 在函数结束时调用 Done 方法，表示该任务完成
	fmt.Printf("Worker %d starting\n", id)
	// 模拟工作
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

// WaitGroupCase 演示了如何使用WaitGroup来等待多个协程完成
// 该函数创建5个worker协程，并等待它们全部执行完毕
func WaitGroupCase() {
	var wg sync.WaitGroup
	// 启动5个worker协程
	for i := 1; i <= 5; i++ {
		wg.Add(1) //   启动前递增计数器（必须在 goroutine 外调用！）
		// 循环了五次，add一共加了 5 ，执行了五次 worker，需要wg.Wait()等五次Done将计数器归零后，才结束等待
		go worker(i, &wg)
	}
	// 等待所有协程完成
	wg.Wait()
	/*
		wg.Wait() 的作用是阻塞当前协程，直到所有通过 wg.Add() 添加的任务都调用了 wg.Done()
		有 wg.Wait() 时，主协程不会等待子协程完成，导致程序可能在子协程完成前就退出
	*/
	fmt.Println("All workers done")
}
