package study_case

import (
	"fmt"
	"time"
)

//goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

/*
goroutine 语法格式：
	go 函数名( 参数列表 )
	例如： go f(x, y, z)
*/

func SayHello() {
	for i := 0; i <= 1000; i++ {
		fmt.Println("Hello Goroutine", i)
		time.Sleep(100 * time.Millisecond) // 睡眠100毫秒
	}
}

func GoroutineCase() {
	go SayHello()
	for i := 0; i <= 5; i++ {
		fmt.Println("Hello Main", i)
		time.Sleep(100 * time.Millisecond) // 睡眠100毫秒 这里如果没有等待，会因为主协程执行完了就结束程序，来不及执行go hello创建的协程
	}

}
