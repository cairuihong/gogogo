package study_case

// 为什么要用锁？
import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup

/*
MutexCase函数启动两个goroutine并发执行add函数，
每个goroutine将全局变量x循环增加5000次。
由于没有同步机制，两个goroutine同时读写x变量会导致数据竞争，
最终x的值会小于10000，证明了并发访问共享资源时需要加锁保护。
*/
func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}
func MutexCase() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
