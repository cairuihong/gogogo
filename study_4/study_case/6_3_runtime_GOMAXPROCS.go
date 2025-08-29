package study_case

/*
Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。

默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）。

***Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。

Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。

Go语言中的操作系统线程和goroutine的关系：

1.一个操作系统线程对应用户态多个goroutine。
2.go程序可以同时使用多个操作系统线程。
3.goroutine和OS线程是多对多的关系，即m:n。


*/
import (
	"runtime"
	"time"
)

//通过将任务分配到不同的CPU逻辑核心上实现并行的效果

func a() {
	for i := 1; i < 10; i++ {
		println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		println("B:", i)
	}

}
func RuntimeGOMAXPROCSCase() {

	runtime.GOMAXPROCS(1)
	// 设置为1，表示程序只能使用一个CPU核心来执行Go代码
	//两个goroutine会在单核上交替执行。
	// 需要时可以设置多个cpu核心运行
	go a()
	go b()
	time.Sleep(time.Second)
}
