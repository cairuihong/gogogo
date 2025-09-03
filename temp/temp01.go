package baiju_order

import (
	"fmt"
	"gogogo/baiju_go_script/public_func"
	"sync"
)

// 批量下单结果结构体
type OrderResult struct {
	Index   int         // 请求序号
	Success bool        // 是否成功
	Error   error       // 错误信息
	Data    interface{} // 响应数据
}

// 批量下单函数
func BatchOrderCommit(count int, concurrency int) {
	// 创建带缓冲的channel用于控制并发数
	requests := make(chan map[string]interface{}, concurrency)
	results := make(chan OrderResult, count)

	var wg sync.WaitGroup

	// 启动固定数量的worker goroutines处理请求
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(requests, results)
		}()
	}

	// 启动一个goroutine收集结果
	go func() {
		wg.Wait()
		close(results)
	}()

	// 发送请求参数到channel
	go func() {
		defer close(requests)
		for i := 0; i < count; i++ {
			orderData := generateOrderData(i) // 生成第i个订单参数
			requests <- orderData
		}
	}()

	// 收集并处理结果
	successCount := 0
	failCount := 0

	for result := range results {
		if result.Success {
			successCount++
			fmt.Printf("订单%d下单成功\n", result.Index)
		} else {
			failCount++
			fmt.Printf("订单%d下单失败: %v\n", result.Index, result.Error)
		}
	}

	fmt.Printf("批量下单完成，成功: %d, 失败: %d\n", successCount, failCount)
}

// worker处理函数
func worker(requests <-chan map[string]interface{}, results chan<- OrderResult) {
	index := 0
	for req := range requests {
		err := public_func.PublicPost(resUrl, req)
		results <- OrderResult{
			Index:   index,
			Success: err == nil,
			Error:   err,
			Data:    nil,
		}
		index++
	}
}

// 生成订单参数函数（示例）
func generateOrderData(index int) map[string]interface{} {
	// 基于原始参数创建新的订单参数
	orderData := make(map[string]interface{})
	for k, v := range resData {
		orderData[k] = v
	}

	// 可以根据需要修改某些字段，例如requestId等
	orderData["requestId"] = fmt.Sprintf("8a39c0c3-4c11-4e24-ac06-4cce81de26%02d", index)
	orderData["certificateNum"] = fmt.Sprintf("MTMwNzA1MjAxNzExMDUwMD%02d", index)

	return orderData
}

// 简化版本的批量下单函数
func SimpleBatchOrder(count int) {
	var wg sync.WaitGroup
	successCount := 0
	failCount := 0
	var mu sync.Mutex

	// 控制并发数，避免同时发起太多请求
	semaphore := make(chan struct{}, 10) // 最多10个并发

	for i := 0; i < count; i++ {
		wg.Add(1)
		semaphore <- struct{}{} // 获取信号量

		go func(index int) {
			defer wg.Done()
			defer func() { <-semaphore }() // 释放信号量

			orderData := generateOrderData(index)
			err := public_func.PublicPost(resUrl, orderData)

			mu.Lock()
			if err == nil {
				successCount++
				fmt.Printf("订单%d下单成功\n", index)
			} else {
				failCount++
				fmt.Printf("订单%d下单失败: %v\n", index, err)
			}
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Printf("批量下单完成，成功: %d, 失败: %d\n", successCount, failCount)
}
