package study_case

import (
	"fmt"
	"time"
)

/*
1.
​​Timer 与 Ticker 的区别​​

​​Timer​​：单次触发，用于延迟或超时控制（如 time.NewTimer(2*time.Second)）。

​​Ticker​​：周期性触发，用于定时任务（如 time.NewTicker(1*time.Second)）。

​​底层共性​​：均基于 runtimeTimer结构体，由 Go 运行时调度器管理
*/
func TimerCase() {
	// 1. timer 的基本使用  定时器的创建、使用和时间获取
	// timer1 := time.NewTimer(2 * time.Second)
	// t1 := time.Now()
	// fmt.Printf("t1:%v\n", t1)
	// t2 := <-timer1.C
	// fmt.Printf("t2:%v\n", t2)

	// 2.验证timer只能响应1次
	/*
		由于没有重置定时器，程序会在第一次触发后阻塞在通道接收操作上，无法继续执行
		原因： time.NewTimer() 创建的定时器只会触发一次，当第一次从 timer2.C 接收到值后，
		 定时器通道不会再有新的值传入，导致程序在第二次循环的 <-timer2.C 处永远阻塞。
	*/
	// timer2 := time.NewTimer(time.Second)
	// for {
	// 	<-timer2.C
	// 	fmt.Println("timer2 触发")
	// }

	// 3.timer实现延时的功能
	//3.1  time.Sleep
	// fmt.Println("start:", time.Now())
	// time.Sleep(time.Second)
	// fmt.Println("start:", time.Now())

	//3.2
	// fmt.Println("timer3:", time.Now())
	// timer3 := time.NewTimer(2 * time.Second)
	// <-timer3.C
	// fmt.Println("timer3:", time.Now())

	// 3.3
	// <-time.After(2 * time.Second)
	// fmt.Println("2秒到")

	// 4 停止定时器  timer4.Stop()
	// timer4 := time.NewTimer(2 * time.Second)
	// go func() {
	// 	<-timer4.C
	// 	fmt.Println("timer4 触发")
	// }()
	// b := timer4.Stop()
	// if b {
	// 	fmt.Println("timer4 已停止")
	// }

	//5 重启定时器
	// timer5 := time.NewTimer(3 * time.Second)
	// timer5.Reset(1 * time.Second)
	// fmt.Println(time.Now())
	// fmt.Println(<-timer5.C)
	// for {
	// }

	// 6. Ticker 的基本使用
	// 6.1 先获取ticker的对象
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	// 6.2 子协程
	go func() {
		for {
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				// 停止定时器
				ticker.Stop()
			}
		}
	}()
	for {

	}
}
