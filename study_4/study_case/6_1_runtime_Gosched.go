package study_case

import (
	"fmt"
	"runtime"
)

/*
runtime.Gosched()
强制主协程让出执行权，让新创建的协程有机会运行
即让出CPU时间片，重新等待安排任务
*/

func RuntimeGoschedCase() {
	// go 创建的匿名函数属于新协程
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	for i := 0; i < 2; i++ {
		runtime.Gosched() // 主协程让出CPU时间片，重新等待安排任务
		fmt.Println("hello")
	}

}
